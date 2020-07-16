package main

import "fmt"

func main() {
	//con el metodo make
	animales := make(map[string]string) //primer argumento es el tipo de dato de la llave y como segundo argumento
	// es el tipo de dato que almacena la clave
	animales["gato"] = "😻" // asignacion de valores
	animales["perro"] = "🐶"

	fmt.Println(animales)
	//sin el metdo make
	frutas := map[string]string{
		"manzana": "🍎",
		"platano": "🍌",
	}
	fmt.Println(frutas)
	//eliminar elemtos del mapa
	delete(frutas, "platano")
	fmt.Println(frutas)

	// obtener los ellemtos del mapa
	fmt.Println(animales["gato"])
	fmt.Println(animales["ardilla"]) //--->  los elemtos que no tienen  clave en el mapa no se imprimen
	// consultar  si uvalor existe en el mapa
	emoji, ok := animales["cocodrilo"]
	fmt.Println(emoji, ok)

	if _, ok := animales["aguila"]; !ok { //usamos l operador blank para que se pueda ejecutar el  programa devido a una variabe que no se esta usando
		animales["agila"] = "🦅"
	}
	fmt.Println(animales)

}
