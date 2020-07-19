package main

import "fmt"

func main() {
	// defer en go
	// nos permite diferir algo  o aplzara la ejecucionde una funcion
	a := 5
	defer fmt.Println("defer:", a) // se aplza la ejeucion de la funcion
	// sin enbargo al momento de plzar la ejecuion a tenia un valor de 5 x eso es que se imprime en ese orden

	a = 10
	fmt.Println(a)

}
