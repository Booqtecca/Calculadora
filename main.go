package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Ingrese el primer número: ")
	fmt.Scanln(&num1)
	fmt.Print("Ingrese el operador (+, -, *, /): ")
	fmt.Scanln(&operator)
	fmt.Print("Ingrese el segundo número: ")
	fmt.Scanln(&num2)

	
	switch operator {
	case "+":
		fmt.Printf("Resultado: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Resultado: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Resultado: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Resultado: %.2f\n", num1/num2)
		} else {
			fmt.Println("Error: No se puede dividir entre cero.")
		}
	default:
		fmt.Println("Operador no válido.")
	}
}
