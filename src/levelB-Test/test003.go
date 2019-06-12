package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var testStr string
		fmt.Scan(&testStr)

		if checkPatValid(testStr) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}

}

func checkPatValid(patStr string) bool {

	if patStr == "PAT" {
		return true
	}

	for _, v := range patStr {
		ch := string(v)

		if !strings.Contains("PAT", ch) {
			return false
		}
	}

	if strings.Contains(patStr, "PAT") {
		n := len(patStr) - 1
		for m := 0; m < len(patStr)/2; m++ {
			n--

			if patStr[m] == patStr[n] && (string(patStr[m]) == "" || string(patStr[m]) == "A") {
				continue
			} else {
				return false
			}

		}

		return true
	}

	return false
}
