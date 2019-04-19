package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"net"
	"sync"
	"time"
)

type myconn struct {
	conn *net.Conn
	i    int
}
type Pool1 struct {
	conns []*myconn

	network string
	address string
	timeout int
	max     int
	mux     sync.Mutex
	active  int
	left    int
}

func (this *Pool1) GetConn(i int) (*myconn, error) {
	this.mux.Lock()
	defer this.mux.Unlock()

	if this.max <= this.active {
		return nil, errors.New("没有连接了")
	}
	if len(this.conns) != 0 { //从里面取一个出来
		conn := this.conns[0]
		this.conns = this.conns[1:]
		this.active++
		//fmt.Println("有连接",this.active,conn.i)
		return conn, nil
	}
	if this.active < this.max {
		conn, err := net.DialTimeout(this.network, this.address, time.Duration(this.timeout)*time.Second)
		if err == nil {
			//this.conns = append(this.conns,&conn)
			this.active++
			//fmt.Println("创建连接",this.active,i)
			conn.SetDeadline(time.Now().Add(time.Duration(this.timeout) * time.Second))
			my := myconn{
				conn: &conn,
				i:    i,
			}
			return &my, nil
		} else {
			return nil, err
		}
	}
	return nil, nil
}
func (this *Pool1) Release(conn *myconn) {
	this.mux.Lock()
	defer this.mux.Unlock()
	this.conns = append(this.conns, conn)
	this.active--
}
func NewPool1(network, address string, max, timeout int) *Pool1 {
	pool := &Pool1{
		network: network,
		address: address,
		max:     max,
		timeout: max,
	}
	pool.conns = []*myconn{}
	//fmt.Println(len(pool.conns))
	//time.Sleep(time.Second*100)
	return pool
}

func main() {
	flag.Parse()
	var pool = NewPool1("tcp", ":6379", 3, 300)

	for i := 1; i < 100; i++ {
		go func(i int) {
			conn, e := pool.GetConn(i)
			time.Sleep(time.Millisecond * 2)
			if e != nil {
				//fmt.Println("读取连接异常",e)
				return
			}
			fmt.Printf("i = %d,myi = %d,\n", i, conn.i)
			defer pool.Release(conn)
		}(i)
		if i%10 == 0 {
			//time.Sleep(time.Second)
		}
	}

	////redis操作
	//v, err := conn.Do("SET", "pool", "test")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(v)
	//v, err = redis.String(conn.Do("GET", "pool"))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(v)

	time.Sleep(time.Second * 1000)
}
