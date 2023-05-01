package main

import "fmt"

func main() {

	hola := "Hola"
	mundo := "mundo"

	fmt.Print(hola)
	fmt.Print(mundo)

	nombre := "Luis"
	edad := 24

	fmt.Println(hola)
	fmt.Println(mundo)

	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)
	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)

	// guardar en una variable
	mensaje := fmt.Sprintf("Hola, %s y su edad es %d", nombre, edad)
	fmt.Println(mensaje)

	// saber el tipo de dato
	fmt.Printf("nombre: %T \n", nombre)

	// insertar dato por teclado
	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)
	fmt.Println("Otro nombre: ", nombre)
}
