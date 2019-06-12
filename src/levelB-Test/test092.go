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

	var N, M int

	var arr = [101][1000]int{}
	fmt.Scan(&N, &M)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < M; i++ {
		s1, _ := reader.ReadString('\n')
		s1 = s1[0 : len(s1)-1]

		arrS1 := strings.Split(s1, " ")

		for j := 0; j < N; j++ {

			arr[i][j], _ = strconv.Atoi(arrS1[j])
		}
	}

	var max int
	var indexs string
	for i := 0; i < N; i++ {
		var sum = 0
		for j := 0; j < M; j++ {
			sum += arr[j][i]
		}
		arr[100][i] = sum

		if sum > max {
			max = sum
		}
	}

	fmt.Println(max)
	for i := 0; i < N; i++ {
		if arr[100][i] == max {
			indexs += strconv.Itoa(i+1) + " "
		}
	}
	fmt.Println(indexs[0 : len(indexs)-1])

}
