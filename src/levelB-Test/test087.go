/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	nMap := make(map[int]int, 2000)

	for i := 1; i <= n; i++ {
		n := i/2 + i/3 + i/5
		if _, exist := nMap[n]; !exist {
			nMap[n] = 1
		}
	}

	fmt.Println(len(nMap))

}
