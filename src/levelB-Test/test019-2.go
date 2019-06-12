/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num int
	var cnt int

	fmt.Scan(&num)

	for num != 6174 && num != 0 || cnt == 0 {
		result, str := parseNum2(num)

		num = result
		fmt.Println(str)
		cnt++
	}

}

func parseNum2(num int) (int, string) {

	var nums = [4]int{}
	var index = 0

	for num > 0 {
		dig := num % 10
		nums[index] = dig
		index++
		num = (num - dig) / 10
	}

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	var n1, n2 int
	var s1, s2 string

	for i := 0; i < 4; i++ {
		n1 *= 10
		n1 += nums[i]

		n2 *= 10
		n2 += nums[3-i]

		s1 += string(nums[i] + '0')
		s2 += string(nums[3-i] + '0')
	}

	s3 := strconv.Itoa(n1 - n2)
	if n1 == n2 {
		s3 = "0000"
	} else if n1-n2 < 1000 {
		s3 = "0" + s3
	}

	return n1 - n2, s1 + " - " + s2 + " = " + s3
}
