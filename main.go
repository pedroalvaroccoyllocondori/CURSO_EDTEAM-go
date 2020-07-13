package main

import "fmt"

func main() {

	// bool,strin, numeric
	var booleano bool = true
	var cadena string = "hola mundo"
	var positivo uint8 = 255
	var bite byte = 255      //alias de uint8
	var negativo rune = -100 //alias de int32
	var caracter rune = 'a'  //alias de unicode code point
	//comillas simples para un solo caracter
	var flotante float64 = 23.56

	//numerioco uint int positivos y negativos respectivamente
	//imprimir el tipo de dato y el valor de la variable
	fmt.Printf("tipo:%T,valor:%v", booleano, booleano) //metodo para imprimir datos
	fmt.Println("")
	fmt.Printf("tipo:%T,valor:%v", cadena, cadena) //metodo para imprimir datos
	fmt.Println("")
	fmt.Printf("tipo:%T,valor:%v", positivo, positivo) //metodo para imprimir datos
	fmt.Println("")
	fmt.Printf("tipo:%T,valor:%v", bite, bite) //metodo para imprimir datos
	fmt.Println("")
	fmt.Printf("tipo:%T,valor:%v", negativo, negativo) //metodo para imprimir datos
	fmt.Println("")
	fmt.Printf("tipo:%T,valor:%v", caracter, caracter) //metodo para imprimir datos
	fmt.Println("")
	fmt.Printf("tipo:%T,valor:%v", flotante, flotante) //metodo para imprimir datos

	// no se pueden ejecutar operaciones con diferente tipo de dato
	resultado := float64(positivo) + flotante // castin para tipo de dato
	// no se cambian los tipos de datos a pesar del casting
	fmt.Println("")
	fmt.Printf("tipo:%T,valor:%v", resultado, resultado) //metodo para imprimir datos

	//operador blank en go
	// en go toda variable declarada tiene que ser usada
	// para dejar variables declaradas que no queremos usar usamos el operador blank
	_ = 230
	var _ string = "prueba"

	//asignacion automatica

	var vacio bool //le asgina el operador nulo encaso de string es un nulo uint es cero
	// en caso de flotantes el operador automaticamente asignado es false
	fmt.Printf("tipo:%T,valor:%v", vacio, vacio) //metodo para imprimir datos

}
