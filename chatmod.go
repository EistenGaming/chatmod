package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	/* Initialize the API to be used */
	client := gptInit()
	/* Initialize the userInput variable. This is what the user enters on the keyboard */
	var userInput string

	/* Loop until the command '/quit' or '/exit' is typed */
	for {

		fmt.Print("> ")
		inputReader := bufio.NewReader(os.Stdin)
		userInput, _ = inputReader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")

		/* Check for the exit condition */
		if userInput == "/quit" || userInput == "/exit" {
			break
		}

		fmt.Println("The question was: " + userInput)
		fmt.Println(gptComplete(client, userInput))

	}

}
