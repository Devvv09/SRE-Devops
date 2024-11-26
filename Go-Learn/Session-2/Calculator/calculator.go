package main

import (
	"fmt"
	"math"
)

func main(){
	
	var choice,num1,num2 float64

	fmt.Println(`
		1 for addition
		2 for substraction
		3 for multiplication
		4 for division
		5 for power
		0 for exit
	
		What is your choice?: 
	`)
	fmt.Scan(&choice)

	fmt.Println("Enter number 1 and number 2")
	fmt.Scan(&num1, &num2)

	switch choice{
		case 1:
			fmt.Println("Result:",add(num1,num2))
		case 2:
			fmt.Println("Result is:",substract(num1,num2))
			
		case 3:
			fmt.Println("Result:",multiply(num1,num2))
			
		case 4:
			fmt.Println("Result is:",division(num1,num2))
			
		case 5:
			fmt.Println("Result is",pow(num1,num2))
			
	}
}

func add(num1,num2 float64) float64 {
	return num1 + num2
}

func substract(num1,num2 float64) float64 {
	return num1 - num2
}

func multiply(num1,num2 float64) float64 {
	return num1 * num2
}

func division(num1,num2 float64) float64 {
	return float64(num1 / num2)
}

func pow(num1,num2 float64) float64 {
	return (math.Pow((num1),(num2)))
}

