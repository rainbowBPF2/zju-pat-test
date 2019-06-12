package main

import (
	"fmt"
)

func main() {
	var A string
	var B int
	var result string

	fmt.Scan(&A, &B)

	var base = 0
	for i := 0; i < len(A); i++ {
		num := int(A[i]) - int('0')

		base *= 10
		base += num

		if base >= B {
			q := base / B
			result += string('0' + q)
			base = base % B
		} else {
			result += string('0')
		}

	}

	result += " " + string('0'+base)
	if result[0] == '0' && result[1] > '0' {
		fmt.Println(result[1:])
	} else {
		fmt.Println(result)
	}
}
