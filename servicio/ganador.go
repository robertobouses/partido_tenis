package servicio

import "fmt"

func (j Jugador) GanadorJugador1(puntuacionJugador1, puntuacionJugador2 string) {
	if puntuacionJugador1 == "GANADOR" {
		fmt.Println("ganador del partido", j.Nombre)
	} else if puntuacionJugador2 == "GANADOR" {

	}
}

func (j Jugador) GanadorJugador2(puntuacionJugador1, puntuacionJugador2 string) {
	if puntuacionJugador2 == "GANADOR" {
		fmt.Println("ganador del partido", j.Nombre)
	} else if puntuacionJugador2 == "GANADOR" {

	}
}

/*package servicio

import "fmt"

func (j Jugador) Ganador(puntuacionJugador1, puntuacionJugador2 string) {
	if puntuacionJugador1 == "GANADOR" {
		fmt.Println("ganador del partido", j.Nombre)
	} else if puntuacionJugador2 == "GANADOR" {
		fmt.Println("ganador del partido", j.Nombre)
	}
}
*/
