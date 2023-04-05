package main

import (
	"fmt"
)

func main() {
	client := gptInit()
	fmt.Println(gptComplete(client, "hello"))
}
