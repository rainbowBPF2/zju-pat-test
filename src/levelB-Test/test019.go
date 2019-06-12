/*
 * Copyright (c) 2019. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package main

import (
	"fmt"
)

func main() {

	var num string
	var numCnt [10]int

	fmt.Scan(&num)

	for i := 0; i < len(num); i++ {
		dig := num[i] - '0'
		numCnt[dig]++

	}

	for i := 0; i < 10; i++ {
		if numCnt[i] > 0 {
			fmt.Printf("%d:%d\n", i, numCnt[i])
		}
	}

}
