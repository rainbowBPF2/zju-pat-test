package basic

import (
	"fmt"
)

func main() {

}

func init() {
	fmt.Println("init method!")
	fmt.Println(mapReduceMethod(100))
	test2()
}

func mapReduceMethod(num int) (res int) {
	mpChan := make(chan int, num)
	var result int
	//map
	for i := 1; i <= num; i++ {
		go func(i int) {
			mpChan <- i * i
		}(i)
	}
	for i := 0; i < num; i++ {
		j := <-mpChan
		result += j
	}

	return result
}

func test2() {

	tch := make(chan int, 1)

	tch <- 3

	fmt.Println(<-tch)

}
