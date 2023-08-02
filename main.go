/*
Escribe un programa que meustre cómo transcurre un juego de tenis y quién lo ha gangado.
El programa recibirá una secuencia formada por P1 Player1 o P2, según quién gane cada puntdo del juego.package partidotenis

-las puntuaciones de un juego son LOVE 15 30 40 deuce, ventaja
-ante la secuencia P1 P1 P2 P2 P1 P2 P1 P1
puntuaciones := [6]string{"love", "15", "30", "40", "AD", "GANADOR"}
*/

package main

import "fmt"

func solicitar() (puntoPara int) {
	fmt.Println("Indica el ganador del punto(1/2)")
	fmt.Scanln(&puntoPara)
	return puntoPara
}

func marcador(puntoPara int, mapaPuntuacion map[int]string) (string, string) {
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

		puntoPara = solicitar()
	}
	return mapaPuntuacion[contador1], mapaPuntuacion[contador2]
}

func main() {
	mapaPuntuacion := map[int]string{0: "love", 1: "15", 2: "30", 3: "40", 4: "GANADOR", -1: "AD"}
	jugador1 := "Rafa Nadal"
	jugador2 := "Novak Djokovic"
	fmt.Println(mapaPuntuacion)

	puntoPara := solicitar()
	puntuacionJugador1, puntuacionJugador2 := marcador(puntoPara, mapaPuntuacion)
	fmt.Printf("Puntuación final: %s - %s\n", puntuacionJugador1, puntuacionJugador2)
	if puntuacionJugador1 == "GANADOR" {
		fmt.Println("ganador del partido", jugador1)
	} else if puntuacionJugador2 == "GANADOR" {
		fmt.Println("ganador del partido", jugador2)
	}
}
