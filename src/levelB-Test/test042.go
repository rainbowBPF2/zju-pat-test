/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var str string

	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	str = strings.ToLower(str)

	cMap := make(map[byte]int, 256)

	for i := 0; i < len(str); i++ {
		ch := str[i]
		cMap[ch]++
	}

	var result byte
	var max int

	for k, v := range cMap {
		if k >= 'a' && k <= 'z' || k >= 'A' && k <= 'Z' {
			if v > max || v == max && k < result {
				max = v
				result = k
			}
		}
	}

	fmt.Println(string(result), max)

}
