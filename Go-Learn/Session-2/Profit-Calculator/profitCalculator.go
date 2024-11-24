package main

import (
	"errors"
	"fmt"
	"os"
)

const fileName = "output.txt"

func main() {

	var taxRate, revenue float64
	var expenses int

	fmt.Print("Please enter your tax rate, revenue, expenses:")
	fmt.Scan(&taxRate, &revenue, &expenses)

	profit, err := findProfit(taxRate, revenue, expenses)

	if err != nil {
		fmt.Println(err)
		panic(1)

	} else {

		strProfit := fmt.Sprintf("Profit is %v\n",profit)
		byteData := []byte(strProfit)
		data, err := reader(fileName)

		if err != nil{
			fmt.Println(err)
		}
		
		data = append(data, byteData...)

		err = writer(fileName,[]byte(data))
		if err != nil{
			fmt.Println("Error:",err)
		}
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

func writer(fileName string, data []byte) error {

	err := os.WriteFile(fileName,data,0700)
	if err != nil{
		return fmt.Errorf("Erorr: %v",err)
	}
	return nil
}

func reader(fileName string)([]byte, error){
	
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return data, nil
}