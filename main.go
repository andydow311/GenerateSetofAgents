package main

import (
	"fmt"
	"os"
)

type agents []string

func main() {
	filePath := os.Args[1]

	agents := getAgents(filePath)

	for _, agent := range agents {
		fmt.Println(agent.label)
	}

}
