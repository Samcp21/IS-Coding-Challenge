package domain

import "gonum.org/v1/gonum/mat"

func RotateMatrix(input [][]float64) [][]float64 {
	if len(input) == 0 || len(input[0]) == 0 {
		return nil
	}

	rows := len(input)
	cols := len(input[0])

	rotated := make([][]float64, cols)
	for i := 0; i < cols; i++ {
		rotated[i] = make([]float64, rows)
		for j := 0; j < rows; j++ {
			rotated[i][j] = input[rows-1-j][i]
		}
	}

	return rotated
}

func QRFactorization(input [][]float64) (q [][]float64, r [][]float64) {
	if len(input) == 0 || len(input[0]) == 0 {
		return nil, nil
	}

	rows := len(input)
	cols := len(input[0])

	flat := make([]float64, 0, rows*cols)
	for _, row := range input {
		flat = append(flat, row...)
	}

	m := mat.NewDense(rows, cols, flat)

	var qr mat.QR
	qr.Factorize(m)

	var qMat mat.Dense
	qr.QTo(&qMat)

	var rMat mat.Dense
	qr.RTo(&rMat)

	return denseToSlice(&qMat), denseToSlice(&rMat)
}

func denseToSlice(m *mat.Dense) [][]float64 {
	rows, cols := m.Dims()
	result := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		result[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = m.At(i, j)
		}
	}
	return result
}
