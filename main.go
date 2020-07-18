package main

import "fmt"

func main() {
	// errores en go

	saludo := hola("alvaro")("ccoyllo")
	fmt.Println(saludo)

}

func hola(nombre string) func(string) string {
	return func(apellido string) string {
		return "hola " + nombre + " " + apellido
	}
}
