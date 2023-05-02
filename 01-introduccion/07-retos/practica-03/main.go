package main

import "fmt"

func ivg(valorVenta float32) [2]float32 {

	var ivg float32
	var precioVenta float32

	ivg = valorVenta * 0.18
	precioVenta = valorVenta + ivg

	return [2]float32{ivg, precioVenta}
}

func main() {

	var valorVenta float32

	fmt.Print("Ingresa el valor de venta: ")
	fmt.Scanln(&valorVenta)

	var valorVentaFunc = ivg
	result := valorVentaFunc(valorVenta)

	fmt.Printf("El IGV es %.2f y el precio de venta es %.2f", result[0], result[1])

}
