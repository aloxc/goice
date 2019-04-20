package pool

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestNewGPool(t *testing.T) {
	// 工厂方法创建连接
	fact := func() (net.Conn, error) { return net.DialTimeout("tcp", ":6379", time.Second*5) }

	// 创建config
	poolConfig := &PoolConfig{
		InitCap: 3,
		MaxCap:  10,
		Factory: fact,
	}
	p, err := NewGPool(poolConfig)
	if err != nil {
		fmt.Println(err)
	}
	// release all connection in gpool
	defer p.Close()
	wg := sync.WaitGroup{}
	var count = 3000
	wg.Add(count)
	// return a connection to gpool
	for i := 0; i < count; i++ {
		time.Sleep(time.Millisecond * 10)
		go func(i int) {
			conn, err := p.Get()
			if err != nil {
				fmt.Println("Get error: ", err)
				return
			}
			set(&conn, i)
			wg.Done()
			defer p.Return(conn)

		}(i)
	}

	wg.Wait()
}
func set(conn *net.Conn, i int) {
	req := MultiBulkMarshal("SET", "a"+strconv.Itoa(i), "abcd-"+strconv.Itoa(i))
	_, err := (*conn).Write([]byte(req))
	if err != nil {
		fmt.Println("set 异常", err)
	}
	p := make([]byte, 1024)
	_, err = (*conn).Read(p)
	if err != nil {
		fmt.Println("set 异常", err)
	}
	_, err = ReadLine(p)
	//fmt.Println("set =" + string(bytes))

	req = MultiBulkMarshal("GET", "a"+strconv.Itoa(i))
	_, err = (*conn).Write([]byte(req))
	if err != nil {
		fmt.Println("get 异常", err)
	}
	p = make([]byte, 1024)
	_, err = (*conn).Read(p)

	if err != nil {
		fmt.Println("get 异常", err)
	}
	var s = strconv.Itoa(i)
	fmt.Println("get =" + string(p[4:(9+len(s))]))

}
func MultiBulkMarshal(args ...string) string {
	var s string
	s = "*"
	s += strconv.Itoa(len(args))
	s += "\r\n"

	// 命令所有参数
	for _, v := range args {
		s += "$"
		s += strconv.Itoa(len(v))
		s += "\r\n"
		s += v
		s += "\r\n"
	}

	return s
}
func ReadLine(p []byte) ([]byte, error) {
	for i := 0; i < len(p); i++ {
		if p[i] == '\r' {
			if p[i+1] != '\n' {
				return []byte{}, errors.New("format error")
			}
			return p[0:i], nil
		}
	}
	return []byte{}, errors.New("format error")
}
