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
	fmt.Println(contentString(expensesList))
}

func contentString(expensesList []float32) string {
	builder := strings.Builder{}

	max, min, total, average := expensesDetails(expensesList)

	fmt.Println(" ----------------- ")
	for i, exp := range expensesList {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", exp))

		if i == len(expensesList) - 1 {
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", total))
			builder.WriteString(fmt.Sprintf("Average: %6.2f\n", average))
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", max))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", min))
		}
	}

	return builder.String()
}

func expensesDetails(expensesList []float32) (max, min, total, average float32) {

	if len(expensesList) == 0 {
		return
	}

	min = expenses.Min(expensesList...)
	max = expenses.Max(expensesList...)
	total = expenses.Sum(expensesList...)
	average = expenses.Average(expensesList...)

	return
}

func Export() {

}