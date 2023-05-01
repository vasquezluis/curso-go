package main

import "fmt"

func main() {

	a := 20
	b := 10

	// suma
	result := a + b
	fmt.Println("suma: ", result)

	// resta
	result = a - b
	fmt.Println("resta: ", result)

	// multiplicacion
	result = a * b
	fmt.Println("multiplicacion: ", result)

	// division
	result = a / b
	fmt.Println("division: ", result)

	// division con casting -> 3.0
	var divi float64 = 3.0 / 2.0
	fmt.Println("division: ", divi)

	// modulo
	result = a % b
	fmt.Println("Modulo: ", result)

}
