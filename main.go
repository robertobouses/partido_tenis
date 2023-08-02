package main

import (
	"fmt"

	"github.com/robertobouses/partido_tenis/comunicacion"
	"github.com/robertobouses/partido_tenis/servicio"
)

func main() {
	mapaPuntuacion := map[int]string{0: "love", 1: "15", 2: "30", 3: "40", 4: "GANADOR", -1: "AD"}
	jugador1 := "Rafa Nadal"
	jugador2 := "Novak Djokovic"
	fmt.Println(mapaPuntuacion)

	puntoPara := comunicacion.Solicitar()
	puntuacionJugador1, puntuacionJugador2 := servicio.Marcador(puntoPara, mapaPuntuacion)
	fmt.Printf("Puntuación final: %s - %s\n", puntuacionJugador1, puntuacionJugador2)
	if puntuacionJugador1 == "GANADOR" {
		fmt.Println("ganador del partido", jugador1)
	} else if puntuacionJugador2 == "GANADOR" {
		fmt.Println("ganador del partido", jugador2)
	}
}
