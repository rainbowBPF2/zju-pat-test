/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import "fmt"

func main() {

	var num [10]int
	var nums [50]int
	var cnt = 0
	for i := 0; i < 10; i++ {
		fmt.Scan(&num[i])
		for j := 0; j < num[i]; j++ {
			nums[cnt] = i
			cnt++
		}
	}

	var result string
	var index int

	for i := 0; i < cnt; i++ {
		if nums[i] > 0 {
			result = string(nums[i] + '0')
			index = i
			break
		}
	}

	for i := 0; i < cnt; i++ {
		if index == i {
			continue
		} else {
			result += string(nums[i] + '0')
		}
	}

	fmt.Println(result)

}
