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

	var num string

	fmt.Scan(&num)

	var active = string(num[0])

	var indexE = strings.Index(num, "E")
	var indexP = strings.Index(num, ".")

	var exp = num[indexE+1:]
	expNum, _ := strconv.Atoi(exp)

	originStr := num[indexP-1 : indexE]

	ignorePoint := false

	if expNum > 0 {

		var floatLen = len(originStr) - 2

		if expNum > floatLen {

			for i := 0; i < expNum-floatLen; i++ {
				originStr += "0"
			}
			ignorePoint = true

		} else if expNum == floatLen {
			ignorePoint = true

		}

		originStr = strings.Replace(originStr, ".", "", -1)

		if !ignorePoint {
			originStr = originStr[0:expNum+1] + "." + originStr[expNum+1:]
		}

	} else {

		leftStr := ""

		for i := 0; i < -expNum; i++ {
			leftStr += "0"
		}

		originStr = leftStr + originStr
		originStr = strings.Replace(originStr, ".", "", -1)
		originStr = string(originStr[0]) + "." + string(originStr[1:])

	}

	if active == "-" {
		originStr = "-" + originStr
	}

	fmt.Println(originStr)
}
