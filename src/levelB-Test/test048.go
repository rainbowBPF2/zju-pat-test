/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
)

func main() {

	var A, B string

	fmt.Scan(&A, &B)

	A = ReverseString(A)
	B = ReverseString(B)

	var result string
	for i, j := 0, 0; i < len(A) || j < len(B); i, j = i+1, j+1 {

		var chA, chB uint8
		if i < len(A) {
			chA = A[i] - '0'
		} else {
			chA = 0
		}

		if j < len(B) {
			chB = B[i] - '0'
		} else {
			chB = 0
		}

		if i%2 == 0 { //odd index num
			ch := (chA + chB) % 13

			if ch >= 10 {
				switch ch {
				case 10:
					result += "J"
				case 11:
					result += "Q"
				case 12:
					result += "K"
				default:

				}
			} else {
				result += string(ch + '0')
			}

		} else {
			var ch uint8
			if chB < chA {
				ch = chB + 10 - chA
			} else {
				ch = chB - chA
			}

			result += string(ch + '0')
		}
	}

	result = ReverseString(result)
	//var start int
	//for i := 0; i < len(result); i++ {
	//	if result[i] != '0' {
	//		start = i
	//		break
	//	}
	//}

	fmt.Println(result)

}

func ReverseString(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
