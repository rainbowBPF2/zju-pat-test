package main

import (
	"fmt"
)

var vCnt [3][2]int
var vA, vB byte

func main() {

	var N int

	fmt.Scan(&N)

	var numA = make([]byte, N)
	var numB = make([]byte, N)
	var result = make([]int, N)
	var cnt0, cnt1, cnt2 int

	for i := 0; i < N; i++ {
		fmt.Scanf("%c %c\n", &numA[i], &numB[i])

		if result[i] = compare(numA[i], numB[i]); result[i] < 0 {
			result[i] = compare(numB[i], numA[i]) * (-1)
		}
		if result[i] > 0 {
			cnt0++
			countVictory(numA[i], 'A')
		} else if result[i] == 0 {
			cnt1++
		} else {
			cnt2++
			countVictory(numB[i], 'B')
		}

	}

	vA = sortArr(vCnt, 0)
	vB = sortArr(vCnt, 1)

	fmt.Println(cnt0, cnt1, cnt2)
	fmt.Println(cnt2, cnt1, cnt0)
	fmt.Println(string(vA), string(vB))

}

func compare(a byte, b byte) int {

	if a == b {
		return 0
	} else if a == 'C' && b == 'J' {
		return 1
	} else if a == 'J' && b == 'B' {
		return 1
	} else if a == 'B' && b == 'C' {
		return 1
	} else {
		return -1
	}

}

func countVictory(b byte, mode byte) {
	var index = 0
	if mode == 'B' {
		index = 1
	}

	switch b {
	case 'B':
		vCnt[0][index]++
	case 'C':
		vCnt[1][index]++
	case 'J':
		vCnt[2][index]++
	}
}
func sortArr(arr [3][2]int, i int) byte {

	max := 0
	pool := []byte{'B', 'C', 'J'}
	var ch byte = ' '

	for j := 0; j < 3; j++ {

		if arr[j][i] > max {
			max = arr[j][i]
			ch = pool[j]
		} else if arr[j][i] == max && ch > pool[j] {
			ch = pool[j]
		}

	}

	return ch
}
