package goroutine_pool

import (
	"fmt"
	"time"
)

//这个是工作线程，处理具体的业务逻辑，将jobs中的任务取出，处理后将处理结果放置在results中。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
func main() {

	//两个channel，一个用来放置工作项，一个用来存放处理结果。
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启三个线程，也就是说线程池中只有3个线程，实际情况下，我们可以根据需要动态增加或减少线程。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 添加9个任务后关闭Channel
	// channel to indicate that's all the work we have.
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	//获取所有的处理结果
	for a := 1; a <= 9; a++ {
		<-results
	}
}