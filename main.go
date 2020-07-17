package main

import "fmt"

func main() {
	// funciones en go
	hola("alvaro", "ccoyllo")

}

func hola(primerNombre string, apellido string) {
	fmt.Printf("hola %s con apellido %s", primerNombre, apellido)
}
