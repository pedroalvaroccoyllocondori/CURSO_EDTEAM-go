package main

import "fmt"

func main() {

	//declaracion de  un array
	var comida [3]string
	comida[0] = "🍕"
	comida[1] = "🍔"
	comida[2] = "🍗"
	fmt.Println(comida)
	//declaracion con asignacion
	frutas := [3]string{"🍇", "🍌", "🍉"}
	fmt.Println(frutas)

}
