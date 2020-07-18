package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// errores en go
	contenido, errores := ioutil.ReadFile("./prueba.txt") //devuelve dos valores
	// el primer valor es el contenido

	if errores != nil {
		fmt.Printf("ocurrio un error: %v", errores)
		return // para finalizar la ejecucion de el programa
	}

	fmt.Println(string(contenido)) //debemos de hacer el castin xq no los devuelve en forma de bites

	contenido1, errores1 := ioutil.ReadFile("./inexistente.txt") //devuelve dos valores
	// el primer valor es el contenido
	if errores1 != nil {
		fmt.Printf("ocurrio un error: %v", errores1)
		return // para finalizar la ejecucion de el programa
	}

	fmt.Println(string(contenido1)) //debemos de hacer el castin xq no los devuelve en forma de bites

}
