package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err := getUserInput("Enter revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expense, err := getUserInput("Enter expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Enter tex rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calcFinancials(revenue, expense, taxRate)

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Earnings after tax: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
	storeResults(ebt, profit, ratio)

}

func getUserInput(infoText string) (float64, error) {
	var val float64
	fmt.Print(infoText)
	fmt.Scan(&val)

	if val <= 0 {
		return 0, errors.New("Input is not positive.")
	}

	return val, nil
}

func calcFinancials(revenue, expense, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expense
	profit = revenue*(1-taxRate/100) - expense
	ratio = ebt / profit
	return
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile("financials.txt", []byte(results), 0644)
}
