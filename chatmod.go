package main

import (
	"fmt"
)

func main() {
	client := gptInit()
	fmt.Println(gptComplete(client, "How are you today? I am radiant!"))
}
