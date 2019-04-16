package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

func work(in chan int, out chan error) {
	select {
	case <-time.After(time.Second * 5):
		out <- errors.New("timeout")
		runtime.Goexit()
	case n := <-in:
		fmt.Printf("receive %d\n", n)
		out <- nil
	}
}

func main() {
	in := make(chan int, 10)
	out := make(chan error, 1)
	//for i := 0; i < 3; i++ {
	go work(in, out)
	if err := <-out; err != nil {
		fmt.Println(err)
	}
	//}
	in <- 1234
	go work(in, out)
	fmt.Println(<-out)

	var i = 0
	for {
		i++
		fmt.Println("i", i)
		time.Sleep(time.Second)
	}
}
