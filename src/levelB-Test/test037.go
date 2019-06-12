/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
)

func main() {

	var pay, afford string

	fmt.Scan(&pay, &afford)

	var arr = [2][3]int{}

	parseNumbers(pay, &arr, 0)
	parseNumbers(afford, &arr, 1)

	pNum := arr[0][0]*17*29 + arr[0][1]*29 + arr[0][2]
	aNum := arr[1][0]*17*29 + arr[1][1]*29 + arr[1][2]

	result := aNum - pNum
	tag := ""
	if result < 0 {
		tag = "-"
		result *= -1
	}

	g := result / (17 * 29)
	s := (result - g*17*29) / 29
	n := result % 29

	fmt.Printf("%s%d.%d.%d", tag, g, s, n)

}

func parseNumbers(pay string, arr *[2][3]int, index int) {

	var num int
	var cnt int
	pay += "."

	for i := 0; i < len(pay); i++ {
		if pay[i] == '.' || i == len(pay)-1 {
			arr[index][cnt] = num
			cnt++
			num = 0
			continue
		}
		ch := int(pay[i] - '0')
		num *= 10
		num += ch
	}
}
