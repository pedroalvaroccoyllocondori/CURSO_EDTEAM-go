package main

import "fmt"

func main() {
	// errores en go
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 200, 300, 4000}
	resultados := filtrar(numeros, func(numeros int) bool {
		if numeros <= 10 {
			return true

		}
		return false

	})
	fmt.Println(resultados)

}
func filtrar(numeros []int, funcion func(int) bool) []int {

	resultado := []int{}
	for _, valor := range numeros {
		if funcion(valor) {
			resultado = append(resultado, valor)

		}
	}
	return resultado

}
