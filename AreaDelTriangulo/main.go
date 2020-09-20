package main

import (
	"fmt"
)

func main() {
	var base float64
	var altura float64
	fmt.Println("Área del Triángulo")
	fmt.Print("Ingrese la base: ")
	fmt.Scan(&base)
	fmt.Scanln()
	fmt.Print("Ingrese la altura: ")
	fmt.Scanf("%f", &altura)

	output := (base * altura) / 2
	fmt.Println("El área del triángulo es: ", output)
}
