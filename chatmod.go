package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	/* Initialize the API to be used */
	client := gptInit()
	/* Initialize the userInput variable. This is what the user enters on the keyboard */
	var userInput string
	userInputHistory := []string{}

	/* A string listing all valid commands */
	var helpString = "\n" +
		"Chatmod is a tool for testing LLM models and other generative systems.\n" +
		"\n" +
		"In most cases, you need to type a command, followed by a prompt.\n" +
		"\n\n" +
		"Examples:" +
		"\n\n" +
		"\t/complete Write a 200 word essay about iron prospecting." +
		"\n\n" +
		"\t/dalleGenerate Create an image of a red children's ball with white dots in front of a yellow wall." +
		"\n\n\n" +
		"Available commands:" +
		"\n\n" +
		"\t/complete [Prompt] \t\tUse the complete function of an LLM. This is stateless, a preceding prompt is not taken into consideration.\n" +
		"\t/dalleGenerate [Prompt] \tGenerate an image using DALL-E. The image is saved as 'image_[Unique ID].png' in the current directory.\n" +
		"\t/exit \t\t\t\tTerminate the app.\n" +
		"\t/help \t\t\t\tShow this help text.\n" +
		"\t/history \t\t\tShow a list of commands entered during the current session.\n" +
		"\t/quit \t\t\t\t(see /exit)\n"

	/* Show the welcome text */
	fmt.Print("\nWelcome to chatmod. Type '/help' for available commands.\n")
	/* Loop until the command '/quit' or '/exit' is typed */
	for {
		/* Show the prompt */
		fmt.Print("> ")

		/* Read the user input */
		inputReader := bufio.NewReader(os.Stdin)
		userInput, _ = inputReader.ReadString('\n')
		userInput = strings.TrimSuffix(userInput, "\n")

		/* Add the last user prompt to the history */
		userInputHistory = append(userInputHistory, userInput)

		/* Extract the command string */
		command := ""
		if userInput[0:1] == "/" {
			var tokens = strings.Fields(userInput)
			if len(tokens[0]) > 0 {
				command = tokens[0]
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
		case "/history":
			history(userInputHistory)
		case "/complete": /* /complete needs to be last as it is the default behavior, which 'falls through' to the default switch statement */
			fallthrough
		default:
			fmt.Println("The prompt was: " + userInput)
			response, err := gptComplete(client, userInput)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(response)
		}

	}

}
