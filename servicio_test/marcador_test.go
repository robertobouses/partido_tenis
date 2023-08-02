package servicio_test

import (
	"testing"

	"github.com/robertobouses/partido_tenis/servicio"
)

type MockSolicitar struct {
	PuntosPara []int
	Index      int
}

func (mock *MockSolicitar) ObtenerPuntoPara() int {
	punto := mock.PuntosPara[mock.Index]
	mock.Index++
	return punto
}

func TestMarcador(t *testing.T) {
	var expected1 string
	var expected2 string
	var result1 string
	var result2 string
	mapaPuntuacion := map[int]string{0: "love", 1: "15", 2: "30", 3: "40", 4: "GANADOR", -1: "AD"}
	mock := &MockSolicitar{
		PuntosPara: []int{1, 1, 1, 1},
		Index:      0,
	}
	punto := mock.ObtenerPuntoPara()
	result1, result2 = servicio.Marcador(punto, mapaPuntuacion)
	expected1 = "GANADOR"
	expected2 = "love"
	if result1 != expected1 || result2 != expected2 {
		t.Errorf("se esperaba %s, %s, se obtuvo %s, %s", expected1, expected2, result1, result2)
	}
}
