package main

import (
	"fmt"
	"math"
)

func displayCalcArt() {
	fmt.Println(`
	_____________________
   |  _________________  |
   | | Calc            | |
   | |_________________| |
   |  ___ ___ ___   ___  |
   | | 7 | 8 | 9 | | + | |
   | |___|___|___| |___| |
   | | 4 | 5 | 6 | | - | |
   | |___|___|___| |___| |
   | | 1 | 2 | 3 | | x | |
   | |___|___|___| |___| |
   | | . | 0 | = | | / | |
   | |___|___|___| |___| |
   |_____________________|
   `)
}

func getUserInput() int {
	var choice int
	fmt.Println("Please select an opperation to perform")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Exponentiate")
	fmt.Println("6. Find the Square Root (n1^n2)")
	fmt.Println("7. Exit")

	for {
		fmt.Print("Enter choice (#): ")
		_, err := fmt.Scan(&choice)
		if err == nil && choice >= 1 && choice <= 7 {
			return choice
		} else {
			fmt.Println("Invaild Input")
		}
	}
}

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	if y == 0 {
		fmt.Println("Error! Division by zero.")
		return 0
	}
	return x / y
}

func exponent(x, y float64) float64 {
	return math.Pow(x, y)
}

func square_rt(x float64) float64 {
	return math.Sqrt(x)
}

func performOperation(choice int) {
	var num1, num2 float64
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	if choice != 6 {
		fmt.Print("Enter second number: ")
		fmt.Scan(&num2)
	}

	switch choice {
	case 1:
		fmt.Println(add(num1, num2))
	case 2:
		fmt.Println(subtract(num1, num2))
	case 3:
		fmt.Println(multiply(num1, num2))
	case 4:
		fmt.Println(divide(num1, num2))
	case 5:
		fmt.Println(exponent(num1, num2))
	case 6:
		fmt.Println(square_rt(num1))
	}
}

func main() {
	for {
		displayCalcArt()
		choice := getUserInput()

		if choice == 7 {
			fmt.Println("Exiting the calculator.")
			break
		}

		performOperation((choice))
		fmt.Println("Press Enter to continue...")
		fmt.Scan()
	}
}
