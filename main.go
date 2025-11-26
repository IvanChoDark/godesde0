package main

import (
	"fmt"
	"github/ivanchodark/godesde0/variables"
)

//"fmt"

func main() {
	//fmt.Println("Hola mundo!!")
	//variables.MostrarEnteros()
	//mismo paquete, archivos diferentes
	//variables.RestoVariables()
	Estado, Texto := variables.ConviertoaTexto(1995)
	fmt.Println(Estado)
	fmt.Println(Texto)
}
