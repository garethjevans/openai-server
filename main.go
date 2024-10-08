package main

import (
	"fmt"

	"github.com/garethjevans/openai-server/server"
)

func main() {
	fmt.Println("OpenAI GPT Mock Server running")
	if err := server.Run(); err != nil {
		panic(err)
	}
}
