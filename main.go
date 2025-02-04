package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	tools.initDirectory("", "asldfasdf")
	for {
		input, _ := reader.ReadString('\n')
		fmt.Println(input)
	}

}
