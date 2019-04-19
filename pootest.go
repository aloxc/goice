package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"net"
	"sync"
	"time"
)

type Pool struct {
	conns []*net.Conn

	network string
	address string
	timeout int
	max     int
	mux     sync.Mutex
	active  int
	left    int
}

func (this *Pool) GetConn(i int) (*net.Conn, error) {
	this.mux.Lock()
	defer this.mux.Unlock()

	if this.max <= this.active {
		return nil, errors.New("没有连接了")
	}
	if len(this.conns) != 0 { //从里面取一个出来
		conn := this.conns[0]
		this.conns = this.conns[1:]
		this.active++
		fmt.Println("有连接", this.active, i, &conn)
		return conn, nil
	}
	if this.active < this.max {
		conn, err := net.DialTimeout(this.network, this.address, time.Duration(this.timeout)*time.Second)
		if err == nil {
			//this.conns = append(this.conns,&conn)
			this.active++
			fmt.Println("创建连接", this.active, i)
			conn.SetDeadline(time.Now().Add(time.Duration(this.timeout) * time.Second))
			return &conn, nil
		} else {
			return nil, err
		}
	}
	return nil, nil
}
func (this *Pool) Release(conn *net.Conn) {
	this.mux.Lock()
	defer this.mux.Unlock()
	this.conns = append(this.conns, conn)
	this.active--
}
func NewPool(network, address string, max, timeout int) *Pool {
	pool := &Pool{
		network: network,
		address: address,
		max:     max,
		timeout: max,
	}
	pool.conns = []*net.Conn{}
	fmt.Println(len(pool.conns))
	//time.Sleep(time.Second*100)
	return pool
}

func main() {
	flag.Parse()
	var pool = NewPool("tcp", ":6379", 4, 300)

	for i := 1; i < 100; i++ {
		go func(i int) {
			conn, e := pool.GetConn(i)
			time.Sleep(time.Millisecond * 20)
			if e != nil {
				//fmt.Println("读取连接异常",e)
			}
			defer pool.Release(conn)
			fmt.Printf("i = %d,conn = %d,\n", i, &conn)
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
