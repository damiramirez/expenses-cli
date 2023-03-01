package main

import (
	"flag"
	"log"
	"strconv"

	"example.com/cli-gastos/commands"
)

func main() {

	var expenses []float32
	var export string

	flag.StringVar(&export, "export", "", "Export the details to .txt")
	flag.Parse()

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

	if export == "" {
		commands.ShowInConsole(expenses)
	} else {
		commands.Export(export, expenses)
	}
}
