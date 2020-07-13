package main

import "fmt"

func main() {
	//operadores  de comparacion: >,<,==,!= ,>=,<=
	fmt.Println(4 > 5)
	fmt.Println(4 != 5)
	fmt.Println(4 == 4)
	//operadores logicos
	var edad = 25
	fmt.Println(edad >= 18 && edad <= 60)

	fmt.Println(edad <= 18 || edad >= 60)
	//operadores unarios :!
	fmt.Println(!(4 != 4))

}
