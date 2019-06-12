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

	var N int
	fmt.Scan(&N)

	var resultMap = make(map[int]int)
	for i := 0; i < N; i++ {
		var class string
		var grade int

		fmt.Scan(&class, &grade)

		index := strings.Index(class, "-")
		cla, _ := strconv.Atoi(class[0:index])

		resultMap[cla] += grade

	}

	var maxCla, maxGrade int

	for k, v := range resultMap {

		if v > maxGrade {
			maxGrade = v
			maxCla = k
		}

	}

	fmt.Println(maxCla, maxGrade)

}
