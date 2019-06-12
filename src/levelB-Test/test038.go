/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var N int

	fmt.Scan(&N)
	scoresMap := make(map[string]int, N/2)

	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	s1 = s1[0 : len(s1)-1]

	s1Arr := strings.Split(s1, " ")

	for _, v := range s1Arr {
		scoresMap[v]++
	}

	s2, _ := reader.ReadString('\n')
	s2 = s2[0 : len(s2)-1]

	s2Arr := strings.Split(s2, " ")
	n, _ := strconv.Atoi(s2Arr[0])

	for i := 1; i <= n; i++ {
		if i == n {
			fmt.Print(scoresMap[s2Arr[i]])
		} else {
			fmt.Print(scoresMap[s2Arr[i]], " ")
		}
	}
}
