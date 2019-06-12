package main

import (
	"fmt"
)

var cacheBit []byte

func main() {

	var n int
	fmt.Scan(&n)

	var cnt = 0
	var lastSu = 0
	cacheBit = make([]byte, n)

	tagSu(n)

	for i := 2; i <= n; i++ {

		if cacheBit[i-1] != '1' && checkSuNum(i) {

			if lastSu > 0 && (i-lastSu == 2) {
				cnt++
			}

			lastSu = i
		}
	}

	fmt.Println(cnt)

}

func checkSuNum(num int) bool {

	if num <= 1 {
		return false
	}

	if num == 2 {
		return true
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func tagSu(upNum int) {

	for i := 2; i <= upNum; i++ {
		if i%2 == 0 {
			cacheBit[i-1] = 1
			continue
		} else if i%3 == 0 {
			cacheBit[i-1] = 1
			continue
		} else if i%5 == 0 {
			cacheBit[i-1] = 1
			continue
		} else if i%7 == 0 {
			cacheBit[i-1] = 1
			continue
		}
	}
}
