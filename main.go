package main

import "fmt"

func main() {
	// funciones  anomimas
	x := func() {
		fmt.Println("hola edtean desde ica")
	}
	x()
	// funcion anomima auto ejecutada
	func(texto string) {
		fmt.Println("soy una funcion autoejecutada", texto)
	}("alvaro")

}
