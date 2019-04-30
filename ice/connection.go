package ice

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/aloxc/goice/utils"
	log "github.com/sirupsen/logrus"
	"github.com/siddontang/go/sync2"
	"net"
	"sync"
	"time"
)


var (
	closedError   = errors.New("连接池正在关闭")
	hasMoreError  = errors.New("没有可用连接了")
	waitMoreError = errors.New("等待请求过多")
)

//func connect(network string, address string, caech chan *connAndError) {
	//log.Info("准备连接服务器",address)
	//var remoteAddress, _ = net.ResolveTCPAddr(network, address)
	//var conn, err = net.DialTCP(network, nil, remoteAddress)
	//if err != nil {
	//	caech <- &connAndError{
	//		error: err,
	//	}
	//	return
	//}
	//c := &Connection{
	//	TCPConn: conn,
	//}
	//c.init()
	//caech <- &connAndError{
	//	Connection: c,
	//}
	//log.Info("已连接服务器",address)
	//return
//}

//like this network:tcp,address :127.0.0.1:1888
func Connect(network string, address string, timeout int) (*theConn, error) {
	var conn, err = net.DialTimeout(network, address, time.Duration(timeout)*time.Second)
	if err != nil {
		return nil, err
	}
	c := &theConn{
		Conn: conn,
	}
	c.init()
	return c, nil
}

type NewConnHook interface {
	hook(conn *net.Conn) error
}

type Pool struct {
	Network     string
	Address     string
	freeConns   []theConn
	usingConns  []theConn
	ConnTimout  int
	MaxConn     int
	idleConn    int
	mux         sync.Mutex
	maxIdle     int
	MaxLifetime time.Duration
	NewConnHook NewConnHook
	closed      sync2.AtomicBool
}
type theConn struct {
	net.Conn
	startTime time.Time
}


func (this *Pool) Get() (*theConn, error) {
	if this.closed.Get() {
		return nil, closedError
	}
	//this.mux.Lock()
	//defer this.mux.Unlock()
	//log.Infof("free %d,using %d,max %d\n", len(this.freeConns), len(this.usingConns), this.MaxConn)
	for {
		if len(this.usingConns) < this.MaxConn { //创建或者取一个空闲的
			if len(this.freeConns) == 0 { //创建一个连接
				log.Info("连接信息：",this.Network,"://",this.Address)
				conn, err := net.DialTimeout(this.Network, this.Address, time.Duration(this.ConnTimout)*time.Second)
				if err != nil {
					log.Info("连接出错",err)
					return nil, err
				}
				if this.NewConnHook != nil {
					err = this.NewConnHook.hook(&conn)
				}
				now := time.Now()
				tConn := &theConn{
					Conn:      conn,
					startTime: now,
				}
				this.usingConns = append(this.usingConns, *tConn)
				return tConn, err
			} else { //取一个连接
				tconn := this.freeConns[0]
				startTime := tconn.startTime
				if startTime.Add(this.MaxLifetime).After(time.Now()) {
					this.usingConns = append(this.usingConns, tconn)
					this.freeConns = this.freeConns[1:]
					return &tconn, nil
				} else { //要关闭该链接了
					tconn.Close()
					log.Info("连接过期了")
					this.freeConns = this.freeConns[1:]
					continue
				}
			}
		} else if len(this.usingConns) == this.MaxConn { //所有都在使用中。只能等待或者返回
			return nil, hasMoreError
		}
	}
}

func (this *Pool) Return(tconn *theConn) {
	//this.mux.Lock()
	//defer this.mux.Unlock()
	this.freeConns = append(this.freeConns, *tconn)
	//log.Info("归还", len(this.freeConns),this.freeConns)

	for index, v := range this.usingConns {
		//log.Info(v == conn)
		if v == *tconn {
			this.usingConns = append(this.usingConns[0:index], this.usingConns[index+1:]...)
			break
		}
	}
}
//关闭连接池
func (this *Pool) Close() {
	this.closed.Set(true)
	log.Info("准备关闭连接池")
	for {
		if len(this.usingConns) > 0 {
			time.Sleep(time.Millisecond * 100)
			log.Info(len(this.usingConns))
			continue
		} else {
			break
		}
	}
	log.Info("没有使用的连接了")
	if len(this.usingConns) == 0 && len(this.freeConns) > 0 {
		wait := sync.WaitGroup{}
		wait.Add(len(this.freeConns))
		for _, conn := range this.freeConns {
			go func() {
				conn.Close()
				wait.Done()
			}()
		}
		wait.Wait()
	}
	this.freeConns = nil
	this.usingConns = nil
	log.Info("关闭完毕")
}
func (this *theConn) init() {
	magic := []byte{0x49, 0x63, 0x65, 0x50}
	var msgType byte = 0
	requestHdr := []byte{
		magic[0],
		magic[1],
		magic[2],
		magic[3],
		GetDefaultProtocolVersion().Major,
		GetDefaultProtocolVersion().Minor,
		GetDefaultProtocalEncodingVersion().Major,
		GetDefaultProtocalEncodingVersion().Minor,
		msgType,
		1} // Compression status. 0：不支持压缩，1：支持压缩，2：已经压缩
	//还要设置压缩位，见requestHdr里面关于压缩位设置。如果压缩的话，第10位设置为2，并且11 12 13 14设置为压缩后的长度。
	var facet string
	rw := bufio.NewReadWriter(bufio.NewReader(this), bufio.NewWriter(this))
	var buf = NewIceBuff(rw)
	var identity Identity
	var context map[string]string
	var mode byte
	//var pmj, pmn, emj, emn, zip, rmsg byte
	//var requestId int
	var size int
	var operator string
	var head, data []byte
	var err error
	buf.Write(requestHdr)           // 10字节
	buf.Write(utils.IntToBytes(69)) //size 10 +4 = 14
	buf.Write(utils.IntToBytes(1))  //requestId 14+4=18
	identity = Identity{
		name: "HelloIce",
	}
	buf.WriteStr(identity.name)     //18+1+8=27
	buf.WriteStr(identity.category) //27+1=28

	if len(facet) == 0 {
		buf.WriteByte(0) //28+1=29
	} else {
		facets := []string{facet}
		buf.WriteStringArray(facets)
	}
	operator = "ice_isA"
	buf.WriteStr(operator) //29+1+7=37
	mode = 1
	buf.WriteByte(mode) //37+1=38
	context = make(map[string]string)
	buf.WriteStringMap(context) //38+1=39
	//数据整形 java 中 BasicStream.endWriteEncaps方法，大约344行，写此后（39位后的）的数据长度，总数据长度减去39
	buf.Write(utils.IntToBytes(30)) //修正数据长度，是总长度减去写完context后的长度 39 +4 = 43

	buf.WriteByte(1) //encoding major 43+1=44
	buf.WriteByte(1) //encoding minor 44+1=45

	buf.WriteStr("::service::HelloService") //45+1+23=69
	buf.Flush()

	head = make([]byte, 14) //先读取头
	size, err = rw.Read(head)
	if err != nil {
		fmt.Println(size)
		panic(err)
	}
	var __Magic = [4]byte{}
	__Magic[0] = head[0]
	__Magic[1] = head[1]
	__Magic[2] = head[2]
	__Magic[3] = head[3]
	//pmj = head[4]
	//pmn = head[5]
	//emj = head[6]
	//emn = head[7]
	//rmsg = head[8]
	//zip = head[9]
	//fmt.Printf("头 size = %d ,requestId = %d ,msg.size = %d \n", size, utils.BytesToInt(head[10:]), utils.BytesToInt(head[10:]))
	//for i, v := range __Magic {
	//	fmt.Printf("__Magic[%d]=%d\n", i, v)
	//}
	//
	//fmt.Printf("协议版本major = %d,minor = %d\n", pmj, pmn)
	//fmt.Printf("编码版本major = %d,minor = %d\n", emj, emn)
	//fmt.Printf("msg = %d\n", rmsg)
	//fmt.Printf("压缩标示 = %d\n", zip)
	//fmt.Printf("数据长度 = %d\n", utils.BytesToInt(head[10:]))

	data = make([]byte, 40) //先读取头
	size, err = rw.Read(data)
	//requestId = utils.BytesToInt(data[0:4])
	//fmt.Printf("请求ID = %d\n", requestId)
	//fmt.Println("================")
}
