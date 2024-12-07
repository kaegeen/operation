package main

import (
	"fmt"
)

// performOperation performs the calculation based on the operator and operands.
func performOperation(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero is not allowed")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}
}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Simple Calculator")
	fmt.Println("=================")
	fmt.Println("Supported operations: +, -, *, /")

	// Read the first number
	fmt.Print("Enter the first number: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Error: Invalid input. Please enter a valid number.")
		return
	}

	// Read the operator
	fmt.Print("Enter an operator (+, -, *, /): ")
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("Error: Invalid operator.")
		return
	}

	// Read the second number
	fmt.Print("Enter the second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Error: Invalid input. Please enter a valid number.")
		return
	}

	// Perform the calculation
	result, err := performOperation(num1, num2, operator)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	// Print the result
	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
