package iteraciones

import "fmt"

func Iterar() {
	/*	var i int
			for { for eterno
				fmt.Println(i)
				i++
			}
		for i < 10 {
				fmt.Println(i)
				i++
			}
	*/
	//conteo normal
	/*for i := 0; i < 10; i++ {
		fmt.Println("hi ", i)
	}*/
	//saltos de 5
	for i := 0; i < 100; i += 5 {
		if i == 6 {
			break //Rompe ciclo
			//	continue vuelve a la condicion ignorado la secuencia de codigo
		}
		fmt.Println("hi ", i)
	}
}
