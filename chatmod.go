package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	client := gptInit()
	var userInput string
	fmt.Print("> ")
	inputReader := bufio.NewReader(os.Stdin)
	userInput, _ = inputReader.ReadString('\n')
	userInput = strings.TrimSuffix(userInput, "\n")

	fmt.Println("The question was: " + userInput)

	fmt.Println(gptComplete(client, userInput))

}
