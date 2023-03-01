package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/cli-gastos/expenses"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {

	fmt.Print("-> ")
	str, err := reader.ReadString('\n')
	if err != nil {
		return "ERROR: ", err
	}

	str = strings.Replace(str, "\n", "", 1)

	return str, nil
}

func ShowInConsole(expensesList []float32) {
	
	builder := strings.Builder{}

	fmt.Println(" ----------------- ")
	for i, exp := range expensesList {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", exp))

		if i == len(expensesList) - 1 {
			fmt.Println(" ----------------- ")
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", expenses.Sum(expensesList...)))
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", expenses.Max(expensesList...)))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", expenses.Min(expensesList...)))
		}
	}

	fmt.Println(builder.String())
}

func Export() {

}