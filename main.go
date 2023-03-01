package main

import (
	"log"
	"strconv"

	"example.com/cli-gastos/commands"
)

func main() {

	var expenses []float32

	for {
		input, err := commands.GetInput()
		if err != nil {
			log.Fatal(err)
		}

		if input == "cls" {
			break
		}

		expense, err := strconv.ParseFloat(input, 32)
		if err != nil {
			continue
		}

		expenses = append(expenses, float32(expense))
	}

	commands.ShowInConsole(expenses)
}
