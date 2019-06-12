/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var huoExs = [12]string{"tam", "hel", "maa", "huh", "tou", "kes", "hei", "elo", "syy", "lok", "mer", "jou"}
var huoNums = [13]string{"tret", "jan", "feb", "mar", "apr", "may", "jun", "jly", "aug", "sep", "oct", "nov", "dec"}
var huoExsMap = make(map[string]int)

func main() {

	var N string
	reader := bufio.NewReader(os.Stdin)

	N, _ = reader.ReadString('\n')
	N = N[0 : len(N)-1]

	numN, _ := strconv.Atoi(N)

	for i := 0; i < len(huoExs); i++ {
		huoExsMap[huoExs[i]] = (i + 1) * 13
	}

	for i := 0; i < len(huoNums); i++ {
		huoExsMap[huoNums[i]] = i
	}

	var result = ""

	for i := 0; i < numN; i++ {
		var origin string

		origin, _ = reader.ReadString('\n')
		origin = origin[0 : len(origin)-1]

		result += parseOrigin(origin) + "\n"

	}

	fmt.Print(result)

}

func parseOrigin(ori string) string {

	result := ""
	head := ori[0]

	if head >= '0' && head <= '9' {
		//from earth

		num, _ := strconv.Atoi(ori)

		n0 := num / 13
		n1 := num % 13

		if num%13 == 0 && num > 0 {
			result = huoExs[num/13-1]
		} else {
			if num > 13 {
				result = huoExs[n0-1] + " " + huoNums[n1]
			} else {
				result = huoNums[n1]
			}
		}

	} else {
		//from mars
		var num int
		chs := strings.Split(ori, " ")

		for i := 0; i < len(chs); i++ {
			n0 := huoExsMap[chs[i]]
			num += n0
		}

		result = strconv.Itoa(num)

	}

	return result
}
