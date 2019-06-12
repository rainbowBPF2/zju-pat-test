package main

import (
	"fmt"
)

func main() {
	var N, M int
	//fmt.Scan(&N,&M)
	_, err := fmt.Scanf("%d,%d", &N, &M)

	if err != nil {
		fmt.Print(err.Error())
	}

	var dataArr = make([]int, N+M)
	var data int

	if M > N {
		M = M % N
	}

	for i := M; i < N+M; i++ {
		fmt.Scan(&data)
		dataArr[i] = data
	}

	for j := 0; j < M; j++ {
		dataArr[j] = dataArr[N+j]
	}

	for i := 0; i < N; i++ {
		if i != N-1 {
			fmt.Print(dataArr[i], " ")
		} else {
			fmt.Print(dataArr[i])
		}
	}

}
