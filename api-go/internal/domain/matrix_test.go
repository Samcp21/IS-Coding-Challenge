package domain

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	input := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [][]float64{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	result := RotateMatrix(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Error en rotación.\nEsperado: %v\nObtenido: %v", expected, result)
	}
}

func TestRotateMatrix_Empty(t *testing.T) {
	var input [][]float64
	result := RotateMatrix(input)

	if len(result) > 0 {
		t.Errorf("Se esperaba una matriz vacía, se obtuvo %v", result)
	}
}

func TestDomain_EdgeCases(t *testing.T) {
	t.Run("RotateMatrix con matriz vacía", func(t *testing.T) {
		emptyMatrix := [][]float64{}
		result := RotateMatrix(emptyMatrix)
		if result != nil {
			t.Error("Se esperaba nil al rotar una matriz vacía")
		}
	})

	t.Run("QRFactorization con matriz vacía", func(t *testing.T) {
		emptyMatrix := [][]float64{}
		q, r := QRFactorization(emptyMatrix)
		if q != nil || r != nil {
			t.Error("Se esperaba nil al factorizar una matriz vacía")
		}
	})
}
func TestDomain_SuccessCases(t *testing.T) {
	t.Run("Rotación de Matriz Correcta", func(t *testing.T) {
		input := [][]float64{
			{1, 2},
			{3, 4},
		}

		rotated := RotateMatrix(input)

		if rotated[0][0] != 3 || rotated[0][1] != 1 || rotated[1][0] != 4 || rotated[1][1] != 2 {
			t.Errorf("La rotación falló. Se obtuvo: %v", rotated)
		}
	})

	t.Run("Factorización QR Correcta", func(t *testing.T) {
		input := [][]float64{
			{12, -51, 4},
			{6, 167, -68},
			{-4, 24, -41},
		}

		q, r := QRFactorization(input)

		if len(q) != 3 || len(q[0]) != 3 {
			t.Errorf("Dimensiones de matriz Q incorrectas. Se obtuvo: %v", q)
		}
		if len(r) != 3 || len(r[0]) != 3 {
			t.Errorf("Dimensiones de matriz R incorrectas. Se obtuvo: %v", r)
		}
	})
}
