package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var input string
	//fmt.Scanf("%s",&input)  //space-delimit , 空格分开

	s := bufio.NewReader(os.Stdin)
	input, _ = s.ReadString('\n')
	input = input[0 : len(input)-1]

	words := strings.Split(input, " ")

	for i := len(words) - 1; i >= 0; i-- {
		if i == 0 {
			fmt.Print(words[i])
		} else {
			fmt.Print(words[i], " ")
		}
	}
}
