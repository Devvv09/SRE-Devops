package main

import "fmt"

func main() {

	var num int
	fmt.Println("Enter num:")
	fmt.Scan(&num)

	if num % 3 == 0 && num % 5 == 0 {
		fmt.Println("Fizz-buzz!")
	} else if num%3 == 0 {
		fmt.Println("Fizz!")
	} else if num%5 == 0 {
		fmt.Println("Buzz!")
	}
}
