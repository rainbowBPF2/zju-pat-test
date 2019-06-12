/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var c1, c2 int
	fmt.Scan(&c1, &c2)

	dur := (c2 - c1) / 100
	if (c2-c1)%100 >= 50 {
		dur++
	}

	var hh, mm, ss int

	if dur >= 3600 {
		hh = dur / 3600
		dur = dur % 3600
	}

	if dur >= 60 {
		mm = dur / 60
		dur = dur % 60
	}

	ss = dur
	fmt.Printf("%s:%s:%s", getTimeStr(hh), getTimeStr(mm), getTimeStr(ss))
}

func getTimeStr(num int) string {
	if num < 10 {
		return "0" + string(num+'0')
	} else {
		return strconv.Itoa(num)
	}
}
