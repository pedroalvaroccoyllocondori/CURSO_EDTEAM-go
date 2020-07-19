package main

import "fmt"

func main() {
	// defer en go
	// nos permite diferir algo  o aplzara la ejecucionde una funcion
	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	//en este caso se agregan a la pila  el  1 despues  el 2 desapues el 3
	// se imprimen del primero al ultimo elemto en la pila

}
