package main

import "fmt" 

type Calculator struct{
    a,b float64
}

func (v Calculator) plus(a,b float64) float64{
    return a + b
} 

func (v Calculator) minus(a,b float64) float64{
    return a - b
} 

func (v Calculator) mult(a,b float64) float64{
    return a * b
} 

func (v Calculator) div(a,b float64) float64{
    return a / b
}

func main() { 
	var a,b float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&b)
	
	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	

	result := 0.0

	switch operator {
	case "+":
		result = Calculator.plus(&a, &b)
	case "-":
		result = Calculator.minus(&a, &b)
	case "*":
		result = Calculator.mult(&a, &b)
	case "/":
		if b != 0 {
			result = Calculator.div(&a, &b)
		} else {
			fmt.Println("Error: Division by zero")
			return
		}
	default:
		fmt.Println("Invalid operator")
		return
	}
}