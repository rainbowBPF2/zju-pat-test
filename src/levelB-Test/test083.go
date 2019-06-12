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

	reader := bufio.NewReader(os.Stdin)
	numStr, _ := reader.ReadString('\n')

	numStr = numStr[0 : len(numStr)-1]

	nums := strings.Split(numStr, " ")

	cardAbs := [10000]int{}

	for i := 0; i < len(nums); i++ {
		num, _ := strconv.Atoi(nums[i])

		abs := num - i - 1

		if abs < 0 {
			abs *= -1
		}

		cardAbs[abs]++

	}

	for i := 9999; i >= 0; i-- {
		if cardAbs[i] > 1 {
			fmt.Printf("%d %d\n", i, cardAbs[i])
		}
	}

}
