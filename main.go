package main

import (
	"fmt"

	"github.com/pedroalvaroccoyllocondori/go/saludar"
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func main() {
	// los modulos nos permiten  administrar las dependencias de nuestros paquetes
	// y controlar las versiones de los mismos
	// creacion de paquetes

	fmt.Println(saludar.Saludoingles())
	fmt.Println(quote.Hello())
	fmt.Println(quoteV3.Concurrency())

	

}
