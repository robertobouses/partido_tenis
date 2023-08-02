package main

import (
	"fmt"

	"github.com/robertobouses/partido_tenis/comunicacion"
	"github.com/robertobouses/partido_tenis/servicio"
)

func main() {
	mapaPuntuacion := map[int]string{0: "love", 1: "15", 2: "30", 3: "40", 4: "GANADOR", -1: "AD"}
	jugador1 := servicio.Jugador{1, "Rafa Nadal", 37}
	jugador2 := servicio.Jugador{2, "Novak Djokovic", 36}
	fmt.Println(mapaPuntuacion)

	puntoPara := comunicacion.Solicitar()
	puntuacionJugador1, puntuacionJugador2 := servicio.Marcador(puntoPara, mapaPuntuacion)
	fmt.Printf("Puntuación final: %s - %s\n", puntuacionJugador1, puntuacionJugador2)
	jugador1.GanadorJugador1(puntuacionJugador1, puntuacionJugador2)

	jugador2.GanadorJugador2(puntuacionJugador1, puntuacionJugador2)
}
