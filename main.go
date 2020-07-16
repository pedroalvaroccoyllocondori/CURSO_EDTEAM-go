package main

import "fmt"

func main() {

	type cursos struct {
		nombre   string
		profesor string
		pais     string
	}

	basedatos := cursos{
		nombre:   "bases de datos",
		profesor: "lola",
		pais:     "peru",
	}

	git := cursos{
		nombre:   "git",
		profesor: "beto",
		pais:     "bolivia",
	}

	dibujo := cursos{"dibujo", "renato", "peru"}

	pintura := cursos{profesor: "alvaro"} // si no se asignan valores automaticamente se  les asigna el opersdor cero o nulo del string

	fmt.Printf("%+v", basedatos)
	fmt.Println("")
	fmt.Printf("%+v", git)
	fmt.Println("")
	fmt.Printf("%+v", dibujo)
	fmt.Println("")
	fmt.Printf("%+v", pintura)

	//acceder a los  campos

	fmt.Printf("%+v", basedatos.profesor)
	fmt.Println("")
	fmt.Printf("%+v", git.pais)
	fmt.Println("")
	fmt.Printf("%+v", dibujo.profesor)
	fmt.Println("")
	fmt.Printf("%+v", pintura.profesor)

	// CREAR PUNTEROS A UNAESTRUCTURA
	fmt.Println("")
	fmt.Println("estructuras con punteros")
	puntero := &basedatos
	(*puntero).profesor = "basededatosprofe"
	puntero.pais = "venezuela"
	fmt.Printf("%+v", basedatos)
	fmt.Println("")
	fmt.Println("%+v", puntero)

}
