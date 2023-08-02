package comunicacion

import "fmt"

func Solicitar() (puntoPara int) {
	fmt.Println("Indica el ganador del punto(1/2)")
	fmt.Scanln(&puntoPara)
	return puntoPara
}
