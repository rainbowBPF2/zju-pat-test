/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
	"math"
)

func main() {

	var N int
	fmt.Scan(&N)

	var max int

	for i := 0; i < N; i++ {

		var a, b int
		fmt.Scan(&a, &b)

		temp := a*a + b*b

		if temp > max {
			max = temp
		}

	}

	var tmp = float64(max)
	var result = math.Sqrt(tmp)

	fmt.Printf("%.2f", result)

}
