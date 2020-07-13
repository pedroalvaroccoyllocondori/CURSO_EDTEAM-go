package main

import "fmt"

func main() {
	fruta := "🍎"
	//almacenar direcciones en memoria
	var direccion *string
	direccion = &fruta // operador de direcciones de memoria '&'

	fmt.Printf("tipo: %T , valor: %s , direccion: %v \n", fruta, fruta, &fruta)
	fmt.Printf("tipo: %T , valor: %v, desreferenciacion: %s", direccion, direccion, *direccion)
}
