package basic

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	const str = "hello goer!"
	fmt.Println(str)
	fmt.Println(runtime.NumCPU())
	go func() {
		fmt.Println("hello golang!")
		time.Sleep(time.Second)
	}()

	fmt.Println(runtime.NumGoroutine())

	var ch chan int = make(chan int, 100)

	for i := 0; i < 100; i++ {
		go setData(i, ch)
	}

	for i := 0; i < 100; i++ {

		fmt.Println(<-ch)
	}

}

func setData(data int, ch chan int) {
	fmt.Println("set data ", data)
	ch <- data
}

func getData(ch chan int) {
	fmt.Println(<-ch)
}
