package main

import "fmt"

func main() {

	comida := [5]string{"🌭", "🍓", "🍋", "🍔", "🍕"}

	frutas := comida[1:3]
	frutas = append(frutas, "🍍", "🍈", "🍐")

	fmt.Println("comidas", comida)
	fmt.Println("frutas", frutas)

	fmt.Println("tamaño frutas", len(frutas))
	fmt.Println("capacidad frutas", cap(frutas))

}
