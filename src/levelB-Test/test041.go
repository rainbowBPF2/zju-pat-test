/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
	"strings"
)

func main() {

	var N, n int
	fmt.Scan(&N)

	idMap := make(map[string]string)

	for i := 0; i < N; i++ {
		var id, test, real string
		fmt.Scan(&id, &test, &real)
		idMap[test] = id + "-" + real
	}

	fmt.Scan(&n)
	result := ""

	for i := 0; i < n; i++ {
		var test string
		fmt.Scan(&test)
		real := idMap[test]
		realArr := strings.Split(real, "-")
		result += realArr[0] + " " + realArr[1] + "\n"
	}

	fmt.Print(result)

}
