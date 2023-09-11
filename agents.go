package main

import (
	"bufio"
	"fmt"
	"os"
)

type agent struct {
	label string
}

func getAgents(filePath string) []agent {

	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var agents []agent

	for fileScanner.Scan() {
		label := fileScanner.Text()
		agent := agent{
			label: label,
		}
		agents = append(agents, agent)
	}

	readFile.Close()
	return agents
}
