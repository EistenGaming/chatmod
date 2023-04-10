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

	/* A string listing all valid commands */
	var helpString = "\n/complete \t\tUse the complete function of an LLM. This is stateless, i.e. you can not reference a preceding prompt.\n/dalleGenerate \t\tGenerate an image using DALL-E. The image is saved as 'example.png' in the current directory.\n/exit \t\t\tTerminate the app.\n/help \t\t\tShow this help text.\n/quit \t\t\t(see /exit)\n"

	/* Loop until the command '/quit' or '/exit' is typed */
	for {

		/* Show the valid commands at startup */

		/* Show the prompt */
		fmt.Print("> ")
		inputReader := bufio.NewReader(os.Stdin)
		userInput, _ = inputReader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")

		/* Extract the command string */
		command := ""
		if userInput[0:1] == "/" {
			var tokens = strings.Fields(userInput)
			if len(tokens[0]) > 0 {
				command = tokens[0]
				fmt.Println("[DEBUG] Selected command: " + command)
				/* remove the command token from the user input as it is not intended to be sent to the API */
				userInput = strings.TrimLeft(userInput, command)
			}
		}

		/* Check for the exit condition */
		if command == "/quit" || command == "/exit" {
			break
		}

		/* Parse the input for commands */
		switch command {
		case "/dalleGenerate":
			dalleGenerate(client, userInput)
		case "/help":
			fmt.Println(helpString)
		case "/complete": /* /complete needs to be last as it is the default behavior, which 'falls through' to the default switch statement */
			fallthrough
		default:
			fmt.Println("The question was: " + userInput)
			fmt.Println(gptComplete(client, userInput))
		}

	}

}
