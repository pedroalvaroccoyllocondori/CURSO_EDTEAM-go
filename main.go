package main

import (
	"fmt"
)

func main() {

	//funcion panic nos permite entara en panico (finalizar la ejecucion s del el programa)
	division(10, 2)
	division(11, 2)
	division(16, 2)
	division(106, 0) // en esta parte la funcion entrae panico
	// nos muestra toda la pila de errores cuando se ejecuta en la terminal

	division(10, 2)

}
func division(dividendo, divisor int) {
	validarDivision(divisor)
	fmt.Println(dividendo / divisor)
}
func validarDivision(divisor int) {
	if divisor == 0 {
		panic("☠️")

	}
}
