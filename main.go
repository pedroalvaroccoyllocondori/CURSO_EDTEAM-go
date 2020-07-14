package main

import "fmt"

func main() {

	comida := [5]string{"🌭", "🍓", "🍋", "🍔", "🍕"}

	frutas := comida[1:3]
	frutas = append(frutas, "🍍", "🍈", "🍐") // si agregamos elemtos se van remplzado en el array original si no hay
	// mas espacio ya no se remplazan el el array
	//si llega a su capacidad maxima crea un nuevo array duplicando
	// la capcidad anterior
	// cuando se sobrepasa la capacidad anterior  pierde referncia al array de donde fue copiado

	fmt.Println("comidas", comida)
	fmt.Println("frutas", frutas)

	fmt.Println("tamaño frutas", len(frutas))    // tamaño del slice
	fmt.Println("capacidad frutas", cap(frutas)) // hace referencia al array original desde donde se empieza a copiar

	verduras := []string{"🌽", "🥕"}
	verduras = append(verduras, "🥒") //crea un nuevo array con el doble de capacidad
	// con referencia al original
	fmt.Println("verduras:", verduras)
	fmt.Println("tamaño verduras", len(verduras)) // tamaño del slice
	fmt.Println("capacidad frutas", cap(verduras))

	// creacion de arrays  con el metodo make
	animales := make([]string, 0, 3) //el metodo make recibe de argumentos el tipo de variable el tamaño y la capacidad

	animales = append(animales, "🐴", "🐶", "😺", "🐢") //crea un nuevo array con el doble de capacidad
	// con referencia al original
	fmt.Println("animales:", animales)
	fmt.Println("tamaño animales", len(animales))    // tamaño del slice
	fmt.Println("capacidad animales", cap(animales)) // se duplica automaticamente

	//operadores nulos en go

	var utiles []string

	fmt.Println("utiles:", utiles)
	fmt.Println("tamaño utiles", len(utiles)) // tamaño del slice
	fmt.Println("capacidad utiles", cap(utiles))
	fmt.Println(utiles == nil) // los valores nulos en go se especifican con el operador nil

	lugares := []string{}

	fmt.Println("lugares:", lugares)
	fmt.Println("tamaño lugares", len(lugares)) // tamaño del slice
	fmt.Println("capacidad lugares", cap(lugares))
	fmt.Println(lugares == nil) // los valores nulos en go se especifican con el operador nil

}
