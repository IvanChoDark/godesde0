package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var tablaSTR int
var texto string

func tablsmultiplicar() string {

	fmt.Println("Ingrese numero: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		tablaSTR, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ERROR")
			panic("el dato ingresado es incorrecto" + err.Error())
		}
		fmt.Println("Tabla del ", tablaSTR)
		for i := 0; i <= 10; i++ {
			//texto += fmt.Sprintf(tablaSTR, " * ", i, "=", tablaSTR*i)

		}

	}
	return texto

}
