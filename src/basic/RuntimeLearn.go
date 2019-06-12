package basic

import (
	"fmt"
	"runtime"
)

func main() {

	go serviceA()
	go serviceB()

	fmt.Println("Main routine over!")
	var str string

	fmt.Println("************", runtime.NumGoroutine(), "***********")
	fmt.Scan(&str)

}

func serviceA() {

	for i := 0; i < 100; i++ {
		fmt.Println("hello world", i)
		runtime.Gosched()
	}
}

func serviceB() {

	for i := 0; i < 100; i++ {
		fmt.Println("*:", i)
		runtime.Gosched()
	}
}
