package main

import "fmt"

func main() {
	//en go solo existeel for

	//for range  maps es el equivalente a un dicionario emn python

	deportes := map[string]string{"basquet": "🏀", "beisbol": "🏈", "futbol": "⚽️"} //en un mapa en go se colocan  tanto el tipode dato de  la clave  asicomo de el valor

	for llave, valor := range deportes {
		fmt.Println("llave:", llave, ",valor:", valor)

	}

}
