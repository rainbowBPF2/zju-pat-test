/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int

	fmt.Scan(&N)

	var nums []float64
	var numStr string

	reader := bufio.NewReader(os.Stdin)
	numStr, _ = reader.ReadString('\n')
	numStr = numStr[0 : len(numStr)-1]

	strs := strings.Split(numStr, " ")

	for i := 0; i < N; i++ {
		num, _ := strconv.ParseFloat(strs[i], 64)
		nums = append(nums, num)
	}
	var result float64
	for i := 1; i <= N; i++ {

		result += nums[i-1] * float64(Cnk(N, i))
		fmt.Println(nums[i-1], Cnk(N, i))
	}

	fmt.Printf("%.2f", result)

}

func Cnk(N int, k int) int {
	return loop(N) / (loop(k) * loop(N-k))
}

func loop(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
