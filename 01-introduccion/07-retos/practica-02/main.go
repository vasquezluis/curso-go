package main

import "fmt"

func division(num1, num2 int) [2]int {

	var cociente int
	var residuo int

	if num1 > num2 {
		cociente = num1 / num2
		residuo = num1 % num2
	} else {
		cociente = num2 / num1
		residuo = num2 % num1
	}

	return [2]int{cociente, residuo}
}

func main() {

	var num1 int
	fmt.Print("Ingrese el primer numero: ")
	fmt.Scanln(&num1)

	var num2 int
	fmt.Print("Ingrese el segundo numero: ")
	fmt.Scanln(&num2)

	var divisionFunc = division
	result := divisionFunc(num1, num2)

	fmt.Printf("El cociente es: %d \n", result[0])
	fmt.Printf("El residuo es: %d", result[1])
}
