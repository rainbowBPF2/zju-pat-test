/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
	"strings"
)

func main() {

	var expect, real string

	fmt.Scan(&expect, &real)

	expect = strings.ToUpper(expect)
	real = strings.ToUpper(real)

	var result = ""
	for i := 0; i < len(expect); i++ {
		ch := string(expect[i])

		if strings.Index(real, ch) == -1 && strings.Index(result, ch) == -1 {
			result += ch
		}
	}

	fmt.Println(result)

}
