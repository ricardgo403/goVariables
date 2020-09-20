package main

import (
	"fmt"
	"math"
)

func main() {
	var radio float64
	fmt.Println("Área del Círculo")
	fmt.Print("Ingrese el radio: ")
	fmt.Scanf("%f", &radio)

	output := math.Pow(radio, float64(2)) * math.Pi
	fmt.Println("El área del círculo es: ", output)
}
