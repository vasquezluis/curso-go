package main

import "fmt"

func saludar(nombre string) {
	// fmt.Println("Hola mundo")

	fmt.Println("Hola", nombre)
}

func sumar(num1, num2 int) int {
	return num1 + num2
}

func main() {
	saludar("Luis")
	fmt.Println(sumar(2, 4))
}
