package main

import "testing"

func TesteSoma(t *testing.T) {

	total := soma(11, 15)

	if total != 31 {

		t.Errorf("resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)

	}

}
