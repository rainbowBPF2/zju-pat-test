package main

import (
	"fmt"
	"strconv"
)

func main() {

	var s1, s2, s3, s4 string
	fmt.Scan(&s1, &s2, &s3, &s4)

	var b1, b2 byte
	var b3 int
	var cnt = 0

	for i := 0; ; i++ {

		if s1[i] >= 'A' && s1[i] <= 'Z' && cnt == 0 {
			if s1[i] == s2[i] {
				cnt++
				if cnt == 1 {
					b1 = s1[i]
				}
				continue
			}
		}

		if cnt == 1 && s1[i] == s2[i] {
			b2 = s1[i]
		}

		if i == len(s1)-1 || i == len(s2)-1 {
			break
		}

	}

	for i := 0; ; i++ {

		if s3[i] >= 'A' && s3[i] <= 'Z' || s3[i] >= 'a' && s3[i] <= 'z' {
			if s3[i] == s4[i] {
				b3 = i
				break
			}
		}

	}

	var v1 string
	var v2 int

	switch b1 {
	case 'A':
		v1 = "MON"
	case 'B':
		v1 = "TUE"
	case 'C':
		v1 = "WED"
	case 'D':
		v1 = "THU"
	case 'E':
		v1 = "FRI"
	case 'F':
		v1 = "SAT"
	case 'G':
		v1 = "SUN"
	}

	if b2 >= 'A' && b2 <= 'Z' {
		v2 = int(b2) - int('A') + 10
	} else {
		v2 = int(b2) - int('0') + 1
	}

	fmt.Printf("%s %s:%s", v1, getStr(v2), getStr(b3))

}

func getStr(num int) string {

	if num < 10 {
		return "0" + strconv.Itoa(num)
	} else {
		return strconv.Itoa(num)
	}

}
