package main

import (
	"fmt"
)

func main() {
	var gradosFahrenheit float64
	fmt.Println("Área del Círculo")
	fmt.Print("Ingrese los grados fahrenheit (°F): ")
	fmt.Scanf("%f", &gradosFahrenheit)
	var output float64 = (float64(5) / 9) * (gradosFahrenheit - 32)
	fmt.Printf("Son %.2f °C\n", output)
}
