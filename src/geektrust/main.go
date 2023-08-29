package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"geektrust/service"
)

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")

		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the input file")

		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var input [][]string
	for scanner.Scan() {
		args := scanner.Text()
		argList := strings.Fields(args)
		input = append(input, argList)
	}
	metroCardSvc := service.NewMetroCardServiceImpl()
	metroCardSvc.MetroCard(input)
}
