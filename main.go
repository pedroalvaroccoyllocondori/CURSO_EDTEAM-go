package main

import "fmt"

func main() {
	emoji := "😺"
	switch emoji {
	case "😺": // no es necesario usar la palabra break
		fmt.Println("es un gato")
	case "🐶", "🐕": //case de forma multiple
		fmt.Println("es un un perro")
	case "🐦", "🦅": // case de forma multiple
		fmt.Println("es un ave")

	default:
		fmt.Println(" no es ni gato ni perro")

	}

	//swit con estructura logica
	fruta := "🍎"
	switch {
	case fruta == "🍎" || fruta == "🍍": // no es necesario usar la palabra break
		fmt.Println("es una fruta")

	case fruta == "🐦" || fruta == "🦅": // case de forma multiple
		fmt.Println("no es una fruta")

	default:
		fmt.Println(" no es ni fruta ni nada de o anterior")

	}
}
