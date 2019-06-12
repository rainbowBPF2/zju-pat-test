/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import "fmt"

func main() {

	var N int

	var a, b, c, d int

	fmt.Scan(&N)

	var sa, sb int

	for i := 0; i < N; i++ {
		fmt.Scan(&a, &b, &c, &d)
		var aSuccess, bSuccess bool
		if b == a+c {
			aSuccess = true
		}

		if d == a+c {
			bSuccess = true
		}

		if aSuccess && !bSuccess {
			sb++
		} else if !aSuccess && bSuccess {
			sa++
		}
	}

	fmt.Println(sa, sb)

}
