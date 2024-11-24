package main

import (
	"errors"
	"fmt"
)

func main() {

	var taxRate, revenue float64
	var expenses int

	fmt.Print("Please enter your taxRate, revenue, expenses:")
	fmt.Scan(&taxRate, &revenue, &expenses)

	profit, err := findProfit(taxRate, revenue, expenses)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(profit)
	}

}

func findProfit(taxRate, revenue float64, expenses int) (float64, error) {
	if taxRate < 0 || revenue < 0 || expenses < 0 {
		return 0.0, errors.New("Error: values cannot be less than zero")
	} else if taxRate == 0 || revenue == 0 || expenses == 0 {
		return 0.0, errors.New("Error: values cannot be zero")
	}

	EBT := revenue - float64(expenses)
	return EBT * (1 - taxRate/100), nil

}
