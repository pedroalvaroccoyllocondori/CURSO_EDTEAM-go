package main

import "fmt"

func main() {
	// funciones variaticas

	fmt.Println(suma(20, 30, 7, 8))

}
func suma(numeros ...int) int {
	resultado := 0
	for _, valor := range numeros {
		resultado += valor
	}
	return resultado
}
