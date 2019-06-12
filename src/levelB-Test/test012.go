package main

import (
	"fmt"
	"strconv"
)

func main() {

	var N int
	fmt.Scan(&N)

	var A = [5]int{}
	var tagA = [5]bool{}
	var cnt = 0
	var cnt2 = 0

	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)

		if num%10 == 0 {
			tagA[0] = true
			A[0] += num
		}

		if num%5 == 1 {
			tagA[1] = true

			if cnt%2 == 0 {
				A[1] += num
			} else {
				A[1] -= num
			}
			cnt++
		} else if num%5 == 2 {
			tagA[2] = true
			A[2]++
		} else if num%5 == 3 {
			tagA[3] = true
			cnt2++
			A[3] += num
		} else if num%5 == 4 {
			tagA[4] = true
			if num > A[4] {
				A[4] = num
			}
		}
	}
	var result = [5]string{}
	for i := 0; i < 5; i++ {
		if tagA[i] == false {
			result[i] = "N"
		} else if i != 3 {
			result[i] = strconv.Itoa(A[i])
		}
	}

	if tagA[3] {
		fmt.Printf("%s %s %s %.1f %s", result[0], result[1], result[2], float32(A[3])/float32(cnt2), result[4])
	} else {
		fmt.Printf("%s %s %s %s %s", result[0], result[1], result[2], "N", result[4])
	}

}
