package main

import "fmt"

func main() {
	emoji := "🌵"
	if emoji == "🌵" {
		fmt.Println("es un cactus")

	} else if emoji == "😃" {
		fmt.Println("es un cactus")
	} else {
		fmt.Println("no es ninguna de las anterioes")
	}
	/// otro tipo de  formar el if en go
	/// con la  condicon en el escope
	if animal := "🐳"; animal == "🌵" {
		fmt.Println("es un cactus")

	} else if animal == "🐳" || emoji == "🐆" {
		fmt.Println("es una ballena")
	} else {
		fmt.Println("no es ninguna de las anterioes")
	}

}
