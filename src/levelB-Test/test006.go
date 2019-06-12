package main

import "fmt"

func main() {

	var num int
	var numSlice []int

	fmt.Scan(&num)
	for num != 0 {
		n := num % 10
		num = (num - n) / 10
		numSlice = append(numSlice, n)
	}

	for i := len(numSlice) - 1; i >= 0; i-- {

		for j := 0; j < numSlice[i]; j++ {
			if i == 2 {
				fmt.Print("B")
			} else if i == 1 {
				fmt.Print("S")
			} else {
				fmt.Print(j + 1)
			}
		}
	}
}
