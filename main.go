package main

import (
	"errors"
	"fmt"
)

func main() {
	// errores en go

	resultado, errorr := division(10, 5)
	if errorr != nil {
		fmt.Printf("ocurrio un error: %v", errorr)
		return // para finalizar la ejecucion de el programa
	}
	fmt.Println("la respuesta es :", resultado)

}

func division(dividendo, divisor int) (resultado int, err error) {

	if divisor == 0 {
		err = errors.New("no puedes dividir entre cero")
		return //retornara valor cero
	}
	resultado = dividendo / divisor
	return //retornara el valor cero de error osea (nil)
}
