package application

import (
	"api-go/internal/domain"
	"api-go/internal/infrastructure/client"
	"errors"
	"fmt"
	"sync"
)

type MatrixService struct {
	nodeClient *client.NodeGRPCClient
}

func NewMatrixService(nodeClient *client.NodeGRPCClient) *MatrixService {
	return &MatrixService{nodeClient: nodeClient}
}

func (s *MatrixService) ProcessMatrix(input [][]float64) (map[string]interface{}, error) {
	if len(input) == 0 || len(input[0]) == 0 {
		return nil, errors.New("la matriz ingresada está vacía o es inválida")
	}

	var q, r, rotated [][]float64

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		q, r = domain.QRFactorization(input)
	}()

	go func() {
		defer wg.Done()
		rotated = domain.RotateMatrix(input)
	}()

	wg.Wait()

	var nodeStats map[string]interface{}

	if s.nodeClient != nil {
		stats, err := s.nodeClient.SendRotatedMatrix(rotated)
		if err != nil {
			return nil, fmt.Errorf("error obteniendo estadísticas de Node: %w", err)
		}
		nodeStats = map[string]interface{}{
			"max_value":   stats.MaxValue,
			"min_value":   stats.MinValue,
			"average":     stats.Average,
			"total_sum":   stats.TotalSum,
			"is_diagonal": stats.IsDiagonal,
		}
	}

	result := map[string]interface{}{
		"factorization_qr": map[string]interface{}{
			"Q": q,
			"R": r,
		},
		"statistics_node": nodeStats,
	}

	return result, nil
}
