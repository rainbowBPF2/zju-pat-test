package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var str string

	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	str = str[0 : len(str)-1]

	arr := strings.Split(str, " ")
	result := make([]int, len(arr))
	resultstr := ""

	for i := 0; i < len(arr); i++ {

		n1, _ := strconv.Atoi(arr[i])

		if i%2 == 0 {
			n2, _ := strconv.Atoi(arr[i+1])
			result[i] = n1 * n2
		} else {

			if n1-1 >= 0 {
				result[i] = n1 - 1
			} else {
				result[i] = 0
				break
			}
		}
	}

	for i := 0; i < len(result); i++ {

		if i%2 == 1 && result[i] == 0 {
			resultstr += strconv.Itoa(result[i])
			break
		} else {
			resultstr += strconv.Itoa(result[i]) + " "
		}
	}

	fmt.Print(strings.TrimSpace(resultstr))
}
