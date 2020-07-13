package main

import "fmt"

func main() {
	//operadores aritmeticos (),*,/,%,+,-
	var a = 4 + 2
	fmt.Println(a)
	//operadores de asignacion =,-=,+=,/=,%=
	var b = 10
	b += 2
	fmt.Println(b)
	// declaracion post-incremento y post-decremento: ++ , --
	// no hay predecremento y predecrecimiento
	// no son una expresion sino una declaracion
	var c = 3
	c--
	fmt.Println(c) //no se puede declarar en el print # ojo

}
