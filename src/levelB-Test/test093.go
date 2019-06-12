/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var A, B string

	reader := bufio.NewReader(os.Stdin)
	A, _ = reader.ReadString('\n')
	B, _ = reader.ReadString('\n')

	C := A[0:len(A)-1] + B[0:len(B)-1]

	chMap := make(map[byte]int)
	result := ""

	for i := 0; i < len(C); i++ {
		ch := C[i]

		if _, exist := chMap[ch]; !exist {
			result += string(ch)
			chMap[ch] = 1
		}
	}

	fmt.Printf("%s", result)

}
