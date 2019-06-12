package main

import "fmt"

func main() {

	var N int
	var ch string

	fmt.Scan(&N, &ch)

	var level = 0

	for 2*level*level-1 <= N {
		level++
	}

	level--

	left := N - (2*level*level - 1)
	innerLevel := level
	var turned = false

	for i := 0; i < 2*level-1; i++ {

		for j := 0; j < (level - innerLevel); j++ {
			fmt.Print(" ")
		}

		for j := 0; j < 2*innerLevel-1; j++ {
			fmt.Print(ch)
		}

		fmt.Println()

		if innerLevel > 1 && !turned {
			innerLevel--
		} else {
			innerLevel++
			turned = true
		}
	}

	fmt.Println(left)

}
