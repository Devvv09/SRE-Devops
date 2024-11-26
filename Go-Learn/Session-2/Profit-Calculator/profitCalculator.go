package main

import (
	"errors"
	"fmt"
	"os"
)

const fileName = "output.txt"

func main() {

	var taxRate, revenue, expenses float64

	fmt.Print("Please enter your tax rate, revenue, expenses:")
	fmt.Scan(&taxRate, &revenue, &expenses)

	profit, err := findProfit(taxRate, revenue, expenses)
	if err != nil {
		panic(fmt.Sprintf("%v", err))

	}
	
	strProfit := fmt.Sprintf("Profit is %v\n", profit)
	byteData := []byte(strProfit)
	data := reader(fileName)

	data = append(data, byteData...)
	writer(fileName, []byte(data))
	
}

func findProfit(taxRate, revenue , expenses float64) (float64, error) {
	if taxRate <= 0 || revenue <= 0 || expenses <= 0 {
		return 0.0, errors.New("Error: values cannot be less than zero")
	} 

	ebt := revenue - expenses
	return ebt * (1 - taxRate/100), nil

}

func writer(fileName string, data []byte) {
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func reader(fileName string) ([]byte) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	return data
}
