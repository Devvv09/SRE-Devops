package main

import (
	"fmt"
	"math"
)

func main(){
	
	var choice,num1,num2 int

	fmt.Println("1 for addition")
	fmt.Println("2 for substraction")
	fmt.Println("3 for multiplication")
	fmt.Println("4 for division")
	fmt.Println("5 for power")

	fmt.Println("Enter your choice!")
	fmt.Scan(&choice)

	fmt.Println("Enter number 1 and number 2")
	fmt.Scan(&num1,&num2)

	switch choice{
		case 1:
			fmt.Println("Result:",add(num1,num2))
			break
		case 2:
			fmt.Println("Result is:",substract(num1,num2))
			break
		case 3:
			fmt.Println("Result:",multiplication(num1,num2))
			break
		case 4:
			fmt.Println("Result is:",division(num1,num2))
			break
		case 5:
			fmt.Println("Result is",pow(num1,num2))
			break
	}
}

func add(num1,num2 int) int {
	return num1 + num2
}

func substract(num1,num2 int) int {
	return num1 - num2
}

func multiplication(num1,num2 int) int {
	return num1 * num2
}

func division(num1,num2 int) float32 {
	return float32(num1 / num2)
}

func pow(num1,num2 int) int {
	return int(math.Pow(float64(num1),float64(num2)))
}

