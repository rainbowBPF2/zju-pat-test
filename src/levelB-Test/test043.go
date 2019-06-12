/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string

	fmt.Scan(&str)

	chMap := make(map[byte]int)

	var cnt = 0

	for i := 0; i < len(str); i++ {
		ch := str[i]

		if strings.Contains("PATest", string(ch)) {
			chMap[ch]++
			cnt++
		}
	}

	var s = "PATest"
	var result = ""

	for cnt > 0 {
		for i := 0; i < len(s); i++ {
			ch := s[i]
			if c, _ := chMap[ch]; c > 0 {

				result += string(ch)
				chMap[ch]--
				cnt--
			}

		}
	}

	fmt.Println(result)

}
