package basic

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	rd := bufio.NewReader(os.Stdin)
	fmt.Println("begin getting data:")

	for {

		line, err := rd.ReadString('\n')
		if err != nil {
			errors.New(err.Error())
		}

		if err == io.EOF {
			log.Print(err.Error())
			break
		}

		if strings.Contains(line, "quit") {
			fmt.Println(line)
			break
		}

		fmt.Println(line)

	}

	yesOrNo()

}

func yesOrNo() {

	var str string

here:
	fmt.Println("Reading new data:")
	fmt.Scan(&str)

	switch str {

	case "yes":
		fmt.Println("continue")
		goto here
	case "no":
		fmt.Println("end here")
	default:
		goto here

	}

}
