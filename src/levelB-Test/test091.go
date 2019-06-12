/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var n int
	fmt.Scan(&n)

	result := []string{}
	for i := 0; i < n; i++ {
		var n int
		fmt.Scan(&n)
		result = append(result, getSelfBaseNum(n))
	}

	for _, v := range result {
		fmt.Println(v)
	}
}

func getSelfBaseNum(num int) string {

	var result = "No"

	for i := 0; i < 10; i++ {
		newNum := num * num * i
		newNumStr := strconv.Itoa(newNum) + ":"

		oldNum := num
		oldNumStr := strconv.Itoa(oldNum) + ":"

		if strings.Contains(newNumStr, oldNumStr) {
			result = fmt.Sprintf("%d %d", i, newNum)
			break
		}

	}

	return result
}
