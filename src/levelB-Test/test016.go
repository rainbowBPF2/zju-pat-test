package main

import "fmt"

func main() {

	var a, da, b, db int
	var pa, pb int

	fmt.Scan(&a, &da, &b, &db)

	pa = getPV(a, da)
	pb = getPV(b, db)

	fmt.Println(pa + pb)

}

func getPV(a int, da int) int {
	var pv int

	for a > 0 {

		num := a % 10
		a = (a - num) / 10

		if num == da {
			pv *= 10
			pv += num
		}

	}
	return pv
}
