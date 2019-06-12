package main

import "fmt"

func main() {

	var N int
	var a, b, c int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&a, &b, &c)

		if a+b > c {
			fmt.Printf("Case #%d: %s\n", i+1, "true")
		} else {
			fmt.Printf("Case #%d: %s\n", i+1, "false")
		}
	}
}
