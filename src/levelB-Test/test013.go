package main

import "fmt"

func main() {

	var M, N, cnt int

	fmt.Scan(&M, &N)
	var result []int

	for i := 1; ; i++ {
		if checkSu(i) {
			cnt++

			if cnt >= M && cnt <= N {
				result = append(result, i)
			}

			if cnt == N {
				break
			}
		}
	}

	for i := 1; i <= len(result); i++ {
		if i%10 == 0 {
			fmt.Println(result[i-1])
		} else {
			if i != len(result) {
				fmt.Print(result[i-1], " ")
			} else {
				fmt.Print(result[i-1])
			}
		}
	}

}

func checkSu(num int) bool {

	if num == 1 {
		return false
	} else if num == 2 {
		return true
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
