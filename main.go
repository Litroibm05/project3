package main

import "fmt"

func main() {

	colores := map[string]string{
		"yellow": "Amarillo",
		"Blue":   "Azul",
		"Red":    "Rojo",
		"Black":  "Negro",
	}

	for k, v := range colores {
		fmt.Println(k, "en español es :", v)
	}
}
