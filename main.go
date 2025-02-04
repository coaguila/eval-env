package main

import (
	"bufio"
	"fmt"
	"os"
)

const PLAY = "/home/pochita/Documents/projects/go-projects/eval-env/playground"

func main() {

	reader := bufio.NewReader(os.Stdin)

	initDirectory(PLAY, "test")
	for {
		input, _ := reader.ReadString('\n')
		fmt.Println(input)
	}

}
