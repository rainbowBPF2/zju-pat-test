package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := ""

	_, err := fmt.Scanf("%s", &num)

	if err != nil {
		panic(err)
	}

	diSum := parseNum(num)
	diStr := parseDi(diSum)
	fmt.Println(diStr[0 : len(diStr)-1])

}

func parseNum(num string) int {
	diSum := 0
	for i := range num {

		di := string(num[i])

		switch di {
		case "1":
			diSum += 1
		case "2":
			diSum += 2
		case "3":
			diSum += 3
		case "4":
			diSum += 4
		case "5":
			diSum += 5
		case "6":
			diSum += 6
		case "7":
			diSum += 7
		case "8":
			diSum += 8
		case "9":
			diSum += 9
		case "0":
			diSum += 0

		}

	}

	return diSum
}

func parseDi(di int) string {

	result := ""
	diStr := strconv.Itoa(di)

	for i := range diStr {

		s := string(diStr[i])

		switch s {
		case "1":
			result += "yi "
		case "2":
			result += "er "
		case "3":
			result += "san "
		case "4":
			result += "si "
		case "5":
			result += "wu "
		case "6":
			result += "liu "
		case "7":
			result += "qi "
		case "8":
			result += "ba "
		case "9":
			result += "jiu "
		case "0":
			result += "ling "
		default:

		}

	}

	return result
}
