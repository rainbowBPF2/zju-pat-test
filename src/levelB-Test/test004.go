package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var n int
	fmt.Scan(&n)

	rd := bufio.NewReader(os.Stdin)

	var personSlice Persons
	for i := 0; i < n; i++ {

		input, _ := rd.ReadString('\n')
		inputStr := strings.Split(input[0:len(input)-1], " ")

		grade, _ := strconv.Atoi(inputStr[2])

		var person = Person{
			name:  inputStr[0],
			id:    inputStr[1],
			grade: grade,
		}

		personSlice = append(personSlice, person)

	}

	sort.Sort(personSlice)

	fmt.Println(personSlice[0].name + " " + personSlice[0].id)
	fmt.Println(personSlice[n-1].name + " " + personSlice[n-1].id)

}

type Person struct {
	name  string
	id    string
	grade int
}

type Persons []Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	return p[i].grade > p[j].grade
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
