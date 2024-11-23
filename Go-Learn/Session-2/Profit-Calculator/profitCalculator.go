package main

import "fmt"

func main(){

	var taxRate,revenue float64
	var expenses int

	fmt.Print("Please enter your taxRate, revenue, expenses:")
	fmt.Scan(&taxRate,&revenue,&expenses)

	fmt.Println("Your profit would be:",findProfit(taxRate,revenue,expenses))

}

func findProfit(taxRate,revenue float64,expenses int) float64 {

	EBT := revenue - float64(expenses)
	return EBT * (1 - taxRate / 100)
	
}