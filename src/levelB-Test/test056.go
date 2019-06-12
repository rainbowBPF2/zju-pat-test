/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import "fmt"

func main() {

	var N int

	fmt.Scan(&N)

	var sli []int

	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		sli = append(sli, num)
	}

	var result int
	for i, v1 := range sli {
		for j, v2 := range sli {
			if i != j {
				num := v1*10 + v2
				result += num
			}

		}
	}

	fmt.Println(result)

}
