package main

import (
	"fmt"
)

func main() {

	var num int
	fmt.Scan(&num)

	cnt := 0
	for num != 1 {
		if num%2 == 0 {
			num = num >> 1
		} else {
			num = (3*num + 1) >> 1
		}

		cnt++
	}

	fmt.Println(cnt)

}
