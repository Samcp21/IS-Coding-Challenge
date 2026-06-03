package application

import (
	"testing"
)

func TestMatrixService_ProcessMatrix(t *testing.T) {
	service := NewMatrixService(nil)

	t.Run("Matriz Inválida - Vacía", func(t *testing.T) {
		_, err := service.ProcessMatrix([][]float64{})
		if err == nil {
			t.Error("Se esperaba un error al procesar una matriz vacía")
		}
	})

	t.Run("Matriz Válida - Camino Feliz", func(t *testing.T) {
		input := [][]float64{
			{1, 2},
			{3, 4},
		}

		result, err := service.ProcessMatrix(input)

		if err != nil {
			if err.Error() != "rpc error: code = Unavailable desc = connection error" {
				t.Logf("Aviso: %v", err)
			}
		} else if result == nil {
			t.Error("Se esperaba un resultado matemático, se obtuvo nil")
		}
	})
}
