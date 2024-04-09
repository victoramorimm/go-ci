package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(112, 10)

	if total != 122 {
		t.Errorf("Resultado da soma é inválido: Resultado esperado: %d, Resultado obtido: %d", 122, total)
	}
}