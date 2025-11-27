package ejercicios

import (
	"strconv"
)

func TestValor(valor string) (int, string) {
	var msn string
	var valorint int
	valorint, _ = strconv.Atoi(valor)

	if valorint > 100 {
		msn = "Es mayor a 100"
	} else {
		msn = "Es menor a 100"
	}
	return valorint, msn
}
