package servicio

import (
	"fmt"

	"github.com/robertobouses/partido_tenis/comunicacion"
)

func Marcador(puntoPara int, mapaPuntuacion map[int]string) (string, string) {
	contador1 := 0
	contador2 := 0

	for {
		if puntoPara == 1 {
			contador1++
		} else if puntoPara == 2 {
			contador2++
		}

		if contador1 == 3 && contador2 == 3 {
			// Deuce
			fmt.Printf("Deuce - Marcador: %s - %s\n", mapaPuntuacion[3], mapaPuntuacion[3])
		} else if contador1 >= 4 && contador2 >= 3 && contador1-contador2 == 1 {
			// Ventaja jugador 1
			fmt.Printf("Ventaja para el jugador 1 - Marcador: %s - %s\n", mapaPuntuacion[-1], mapaPuntuacion[3])
		} else if contador2 >= 4 && contador1 >= 3 && contador2-contador1 == 1 {
			// Ventaja jugador 2
			fmt.Printf("Ventaja para el jugador 2 - Marcador: %s - %s\n", mapaPuntuacion[3], mapaPuntuacion[-1])
		} else if contador1 >= 4 && contador1-contador2 >= 2 {
			// Ganador jugador 1
			fmt.Println("Ganador del partido: Jugador 1")
			return mapaPuntuacion[4], mapaPuntuacion[3]
		} else if contador2 >= 4 && contador2-contador1 >= 2 {
			// Ganador jugador 2
			fmt.Println("Ganador del partido: Jugador 2")
			return mapaPuntuacion[3], mapaPuntuacion[4]
		} else if contador1 >= 3 && contador2 >= 3 && contador1 == contador2 {
			fmt.Printf("Deuce - Marcador: %s - %s\n", mapaPuntuacion[3], mapaPuntuacion[3])

		} else {
			fmt.Printf("Marcador: %s - %s\n", mapaPuntuacion[contador1], mapaPuntuacion[contador2])

		}

		puntoPara = comunicacion.Solicitar()
	}
	return mapaPuntuacion[contador1], mapaPuntuacion[contador2]
}
