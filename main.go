package main

import "fmt"

func main() {
	//en go solo existeel for

	//for range slice funciona npara iterar slices

	numeros := []uint8{2, 4, 6, 8, 10}

	for indice, valor := range numeros {
		fmt.Println("indice", indice, "valor", valor)
	}

	nums := []uint8{2, 4, 6, 8, 10}
	for _, valor := range nums { // no cambia el valor ya que  se esta utilizando una opia del elemeto
		valor *= 5
	}
	fmt.Println(nums)

	for indice, _ := range nums {
		nums[indice] *= 5
	}
	fmt.Println(nums)

}
