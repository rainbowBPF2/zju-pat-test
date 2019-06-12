package basic

import (
	"fmt"
	"runtime"
)

func main() {

	testCh := make(chan int)
	for i := 0; i < 1e7; i++ {
		go newRoutine(testCh, i)
		if i%100 == 0 {
			fmt.Println(runtime.NumGoroutine())
		}
	}

	for i := 0; i < 1e7; i++ {

		fmt.Println(<-testCh)
	}
}

func newRoutine(ch chan int, data int) {
	ch <- data
}
