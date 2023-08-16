package main

import "fmt" 

//определение структуры
type Calculator struct{
    mem float64
}


func (v Calculator) plus(a float64) float64{
     v.mem = a + v.mem
	 return v.mem
} 

func (v Calculator) minus(a float64) float64{
    v.mem = v.mem - a
	return v.mem
} 

func (v Calculator) mult(a float64) float64{
    v.mem = a * v.mem
	return v.mem
} 

func (v Calculator) div(a float64) float64{
    v.mem = v.mem / a
	return v.mem
}

func (v Calculator) clear() {
    v.mem = 0
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

	calc := Calculator{a}
	
	switch operator {
	case "+":
		result = calc.plus(b)
	case "-":
		result = calc.minus(b)
	case "*":
		result = calc.mult(b)
	case "/":
		if b != 0 {
			result = calc.div(b)
		} else {
			fmt.Println("Error: Division by zero")
			return
		}
	default:
		fmt.Println("Invalid operator")
		return
	}

	fmt.Println(result)
}