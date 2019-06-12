/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import "fmt"

func main() {

	var s0, s1 string

	fmt.Scan(&s0, &s1)

	s0Map := make(map[byte]int, 1000)
	s1Map := make(map[byte]int, 1000)

	for i := 0; i < len(s0); i++ {
		ch := s0[i]
		s0Map[ch]++
	}

	for i := 0; i < len(s1); i++ {
		ch := s1[i]
		s1Map[ch]++
	}

	var result = "Yes"
	var left int

	for k, v := range s1Map {
		if v0, exist := s0Map[k]; exist && v0 >= v {
			continue
		} else {
			result = "No"
			left += v - v0
		}
	}

	if result == "Yes" {
		fmt.Println("Yes", len(s0)-len(s1))
	} else {
		fmt.Println("No", left)
	}

}
