/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import "fmt"

func main() {

	var N int
	var ch string

	fmt.Scan(&N, &ch)
	line := (N + 1) / 2

	for i := 0; i < line; i++ {
		if i == 0 || i == line-1 {
			for i := 0; i < N; i++ {
				fmt.Print(ch)
			}
		} else {
			for i := 0; i < N; i++ {
				if i == 0 || i == N-1 {
					fmt.Print(ch)
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}
