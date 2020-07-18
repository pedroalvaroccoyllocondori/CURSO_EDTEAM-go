package main

import (
	"fmt"
	"strings"
)

func main() {
	// funciones en go
	texto := "AlVarO"
	minuscula, mayuscula := convertir(texto)
	fmt.Println("nimuscula:", minuscula, "mayuscula:", mayuscula)

}

// funciones con  multiples valores
func convertir(texto string) (string, string) { // funciones con multiples retornos
	minusculas := strings.ToLower(texto)
	mayusculas := strings.ToUpper(texto)
	return minusculas, mayusculas

}
