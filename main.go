package main

import "fmt"

func main() {
	// defer en go
	// nos permite diferir algo  o aplzara la ejecucionde una funcion
	defer fmt.Println(3) // se ejecuta al finalizar todas la lineas

	fmt.Println(1)
	fmt.Println(2)

}
