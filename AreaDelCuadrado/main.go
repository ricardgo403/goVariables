package main

import (
	"fmt"
	"math"
)

func main() {
	var lado float64
	fmt.Println("Área del Cuadrado")
	fmt.Print("Ingrese el lado: ")
	fmt.Scan(&lado)

	output := math.Pow(lado, float64(2))
	fmt.Println("El área del cuadrado es: ", output)
}
