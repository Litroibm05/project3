package main

import "fmt"

func main() {

	var (
		nombre string
		edad   int
		pi     float64
	)

	fmt.Print("Ingrese su Nombre: ")
	fmt.Scanln(&nombre)

	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&edad)

	fmt.Print("Ingrese valor de Pi: ")
	fmt.Scanln(&pi)

	fmt.Printf("Nombre: %s Edad: %d \nValor de Pi: %f", nombre, edad, pi)

}
