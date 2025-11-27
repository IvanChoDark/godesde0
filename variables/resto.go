package variables

import (
	"fmt"     // para impresiones
	"strconv" // conversiones de datos
	"time"    // para fechas
)

// vaiable Global ya que inicia con mayuscula, para  el que importe este paquete
var Nombre string

// vaiable local ya que inicia con minuscula,
var nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time // para variables con fechas hay que importar libreria time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 17.66
	Fecha = time.Now()
	println(Nombre)
	println(Estado)
	println(Sueldo)
	fmt.Println(Fecha)
}

// (parm entrada)(valores de salida)
func ConviertoaTexto(numeri int) (bool, string) {
	texto := strconv.Itoa(numeri)
	return true, texto
}
