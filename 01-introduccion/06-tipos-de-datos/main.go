// tipos de datos

// Tipos de datos básicos:

//     bool: representa un valor booleano (true o false).
//     string: representa una cadena de texto.
//     int: representa un número entero.
//     float32 y float64: representan números en coma flotante.

// Tipos de datos compuestos:

//     array: representa una colección de elementos del mismo tipo.
//     slice: es similar a un array, pero su tamaño puede ser modificado dinámicamente.
//     map: representa una colección de pares clave-valor.
//     struct: representa un conjunto de campos con diferentes tipos de datos.

// Tipos de datos de referencia:

//     pointer: representa la dirección de memoria de una variable.
//     function: representa una función.
//     interface: define un conjunto de métodos que una estructura debe implementar.

// Tipos de datos especiales:

//     chan: se utiliza para la comunicación entre goroutines.
//     complex64 y complex128: representan números complejos.

package main

import "fmt"

func main() {
	// -----------> Declarar una variable booleana <-----------
	var esMayorDeEdad bool = true
	// Imprimir el valor de la variable
	fmt.Println("Es mayor de edad:", esMayorDeEdad)

	// -----------> Declarar una variable de tipo string <-----------
	var mensaje string = "Hola mundo"
	// Imprimir el valor de la variable
	fmt.Println(mensaje)

	// -----------> Declarar una variable de tipo entero <-----------
	var edad int = 28
	// Imprimir el valor de la variable
	fmt.Println("Edad:", edad)

	// -----------> Declarar una variable de tipo float32 <-----------
	var altura float32 = 1.75
	// Imprimir el valor de la variable
	fmt.Println("Altura:", altura)

	// -----------> Declarar una variable de tipo float64 <-----------
	var pi float64 = 3.1415926
	// Imprimir el valor de la variable
	fmt.Println("Pi:", pi)

	// -----------> Declarar un array de tipo entero <-----------
	var numeros [5]int = [5]int{1, 2, 3, 4, 5}
	// Imprimir el valor de un elemento del array
	fmt.Println("Posicion 2:", numeros[2])

	// -----------> Declarar un slice de tipo entero <-----------
	var numeros2 []int = []int{1, 2, 3, 4, 5}
	// Agregar un nuevo elemento al slice
	numeros2 = append(numeros2, 6)
	// Imprimir el valor de un elemento del slice
	fmt.Println("Posicion 5:", numeros2[5])

	// -----------> Declarar un map con claves de tipo cadena y valores de tipo entero <-----------
	var edades map[string]int = map[string]int{"Juan": 28, "Ana": 30}
	// Imprimir el valor asociado a una clave en el map
	fmt.Println("Edad de Ana:", edades["Ana"])

	// -----------> Declarar una estructura con diferentes campos <-----------
	type Persona struct {
		nombre string
		edad   int
		altura float64
	}

	// -----------> Declarar una variable de tipo Persona <-----------
	var persona1 Persona = Persona{"Juan", 28, 1.75}
	// Imprimir el valor de un campo de la estructura
	fmt.Println("Edad de", persona1.nombre, ":", persona1.edad)

	// -----------> Declarar una variable de tipo entero <-----------
	var edad2 int = 28
	// Declarar un puntero a la variable edad
	var puntero *int = &edad2
	// Imprimir el valor de la variable edad
	fmt.Println("Edad:", edad2)
	// Imprimir el valor de la variable a través del puntero
	fmt.Println("Edad a través del puntero:", *puntero)
	// Actualizar el valor de la variable a través del puntero
	*puntero = 29
	// Imprimir el nuevo valor de la variable
	fmt.Println("Nueva edad:", edad2)

	// -----------> Asignar la función a una variable <-----------
	var funcionSuma = Suma
	// Llamar a la función a través de la variable
	resultado := funcionSuma(1, 2)
	// Imprimir el resultado
	fmt.Println("Suma:", resultado)
}

// -----------> Declarar una función que suma dos números <-----------
func Suma(a, b int) int {
	return a + b
}
