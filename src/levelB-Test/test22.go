/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
)

func main() {

	var A, B, D int

	fmt.Scan(&A, &B, &D)

	sum := A + B

	var resultStr = ""

	for sum > 0 {
		num0 := sum % D
		sum = sum / D
		resultStr = string(num0+'0') + resultStr
	}

	if A+B == 0 {
		fmt.Println(0)
	} else {
		fmt.Print(resultStr)
	}
}
