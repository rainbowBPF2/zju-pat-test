/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var A, B int

	fmt.Scan(&A, &B)

	C := A * B

	strc := strconv.Itoa(C)

	result := ""
	for i := len(strc) - 1; i >= 0; i-- {
		result += string(strc[i])
	}

	var index = 0
	for i := 0; i < len(result); i++ {
		if result[i] != '0' {
			index = i
			break
		}
	}

	fmt.Println(result[index:])

}
