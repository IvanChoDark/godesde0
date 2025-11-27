package main

import (
	"fmt"
	//"github/ivanchodark/godesde0/variables"
	"github/ivanchodark/godesde0/ejercicios"
	//"runtime" // informacion del equipo en que se esta corriendo el app
)

//"fmt"

func main() {
	/*
		//fmt.Println("Hola mundo!!")
		//variables.MostrarEnteros()
		//mismo paquete, archivos diferentes
		//variables.RestoVariables()
		Estado, Texto := variables.ConviertoaTexto(1995)
		fmt.Println(Estado)
		fmt.Println(Texto)
		//os := runtime.GOOS
		if os := runtime.GOOS; os == "Linux." || os == "OS X." {
			fmt.Println("esto no es Windows, es: ", os)
		} else {
			fmt.Println("esto es Windows")
		}

		switch os := runtime.GOOS; os {
		case "Linux. ":
			fmt.Println("Esto es linux")
		case "darwin":
			fmt.Println("Esto es darwin")
		default:
			fmt.Printf("%s \n", os)
		}
	*/
	fmt.Println(ejercicios.TestValor("2000"))
}
