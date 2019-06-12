/*
 * Copyright (c) 2019 Jack-released.
 */

package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var wg sync.WaitGroup
var wg1 sync.WaitGroup

var cnt int

func main() {

	var str string

	fmt.Scan(&str)

	for i := 0; i < len(str); i++ {
		ch0 := str[i]
		if ch0 == 'P' {
			wg.Add(1)
			go findP(i, str)
		}
	}

	wg1.Wait()
	wg.Wait()

	fmt.Println(cnt)

}

func findP(i int, str string) {

	for j := i + 1; j < len(str); j++ {
		ch1 := str[j]

		if ch1 == 'A' {
			wg1.Add(1)
			go findT(j, str, &cnt)

		}
	}
	wg.Done()
}

func findT(j int, str string, cnt *int) {
	for k := j + 1; k < len(str); k++ {
		ch2 := str[k]
		if ch2 == 'T' {
			lock.Lock()

			*cnt++
			if *cnt > 1000000007 {
				*cnt = *cnt % 1000000007
			}
			lock.Unlock()

		}
	}

	wg1.Done()

}
