package main

import (
	"fmt"
	"sort"
)

func main() {

	var n int
	fmt.Scan(&n)

	numMap := make(map[int]int)

	//reading data
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)

		numMap[num] = num
	}

	//find key num
	for k := range numMap {

		dealNumMap(k, numMap)
		fmt.Println("current map:", numMap)
	}

	var keys SliceArray
	for k := range numMap {
		keys = append(keys, k)
	}
	sort.Sort(keys)

	for i, v := range keys {
		if i < len(keys)-1 {
			fmt.Print(v, " ")
		} else {
			fmt.Print(v)
		}
	}
}

func dealNumMap(num int, numMap map[int]int) {

	for num != 1 {
		if num%2 == 0 {
			num = num >> 1
		} else {
			num = (3*num + 1) >> 1
		}

		delete(numMap, num)
	}
}

type SliceArray []int

func (m SliceArray) Len() int {
	return len(m)
}

func (m SliceArray) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m SliceArray) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
