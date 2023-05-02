package main

import "fmt"

func suma(num1, num2 int) int {
	return num1 + num2
}

func main() {

	var num1 int
	fmt.Print("Ingresa el primer numero: ")
	fmt.Scanln(&num1)

	var num2 int
	fmt.Print("Ingresa el primer numero: ")
	fmt.Scanln(&num2)

	var sumaFunc = suma
	result := sumaFunc(num1, num2)

	fmt.Printf("La suma de %d y de %d es: %d", num1, num2, result)
}
