package files

import (
	"bufio"
	"fmt"
	"github/ivanchodark/godesde0/ejercicios"
	"os"
)

var filename string = "./files/txt/Tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo")
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(file string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo" + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		println("> " + registro)
	}
	archivo.Close()
}
