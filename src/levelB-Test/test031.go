/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import "fmt"

func main() {
	var N int
	var weights = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var mapCodes = [11]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	var result = ""

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		var id string
		fmt.Scan(&id)

		var sum = 0
		var skip = false

		for j := 0; j < 17; j++ {
			ch := id[j]
			if ch < '0' || ch > '9' {
				result += id + "\n"
				skip = true
				break
			}

			sum += int(ch-'0') * weights[j]
		}

		if !skip {
			sum %= 11
			mCode := mapCodes[sum]
			if mCode != string(id[17]) {
				result += id + "\n"
			}

		}

	}

	if result == "" {
		fmt.Println("All passed")
	} else {
		fmt.Print(result)
	}
}
