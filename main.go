package main

import "fmt"

func main() {

	//los slices no poseen datos , son apuntadores a un array
	conjunto := [5]string{"🐎", "😸", "🦃", "🐢", "🐄"}
	//slice el ultimo elemeto es excluyente
	//los slices toman partes del array
	animales := conjunto[0:3]
	voladores := conjunto[3:5] // si no ponemos el segundo parametro se rellena con el ultimo .ejemplo= [3:]
	voladores[0] = "🦋"
	fmt.Println(conjunto)
	fmt.Println(animales)
	fmt.Println(voladores)
	fmt.Println(voladores[0])
	fmt.Println(conjunto[:]) // si no especificamos el numero se autoreellena con todos los elemtos

}
