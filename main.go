package main

import (
	"errors"
	"fmt"
)

func main() {
	// errores en go

	resultado, errorr := division(10, 0)
	if errorr != nil {
		fmt.Printf("ocurrio un error: %v", errorr)
		return // para finalizar la ejecucion de el programa
	}
	fmt.Println("la respuesta es :", resultado)

}

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("no puedes dividir entre cero") //devolvemos cero  x q es un valor que no se utiliza
	}
	return dividendo / divisor, nil //devolvemos nil xq es su valor nulo de error
}
