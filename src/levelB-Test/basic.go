/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import "fmt"

func main() {

	fmt.Println(Cnk1(4, 4))

}

func Cnk1(N int, k int) int {
	return Loop(N) / (Loop(k) * Loop(N-k))
}

func Loop(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
