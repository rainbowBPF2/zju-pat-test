/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	var arr = []int{}
	var arrMap = make(map[int]int)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		sum := getSum(num)
		arrMap[sum] = 1

	}

	for k := range arrMap {
		arr = append(arr, k)
	}

	sort.Ints(arr)
	result := ""

	for i := 0; i < len(arr); i++ {
		result += strconv.Itoa(arr[i]) + " "
	}

	fmt.Println(len(arr))
	fmt.Println(result[0 : len(result)-1])

}

func getSum(num int) int {

	var sum int

	for num > 0 {
		n := num % 10
		num = (num - n) / 10
		sum += n
	}

	return sum
}
