package main

import "fmt"

func main() {
	//en go solo existeel for

	//for range  en tipode datos string

	hola := "hola a todos"

	for _, valor := range hola {
		fmt.Println(valor) //si defrente imprimins  los valores de las letras no dan los bites de cada letra

	}
	for _, valor := range hola {
		fmt.Println(string(valor)) //para obtener los valores debemos de  hacer un castin  a los bites de cada letra
	}

}
