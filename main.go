package main

import (
	"fmt"
	"os"
)

func main() {

	//aplicaciones de uso
	// limpiar recursos , cerrar archivos,cerrar conexiones de red, cerar conexiones de base de datos
	archivo, errorr := os.Create("prueba2.txt")

	if errorr != nil {
		fmt.Printf("ocurio un error al crear: %v", errorr)
		return
	}
	defer archivo.Close()                                    // siempre se ejecuta xq esta en la pila del defer
	_, errorr = archivo.Write([]byte("hola alvaro ccoyllo")) // develve una variable n (numero de bytes sritos y un error)

	if errorr != nil {
		//archivo.Close() // debemos cerrar el arivo para poder limpiar los recursos
		fmt.Printf("ocurio un error al escribir : %v", errorr)
		return
	}

}
