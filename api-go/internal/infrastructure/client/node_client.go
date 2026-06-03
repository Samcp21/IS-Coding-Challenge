package client

import (
	"api-go/internal/infrastructure/client/pb"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NodeGRPCClient struct {
	client pb.MatrixServiceClient
}

func NewNodeGRPCClient(address string) (*NodeGRPCClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewMatrixServiceClient(conn)
	return &NodeGRPCClient{client: client}, nil
}

func (c *NodeGRPCClient) SendRotatedMatrix(matrix [][]float64) (*pb.StatsResponse, error) {
	var pbRows []*pb.MatrixRequest_Row
	for _, row := range matrix {
		pbRows = append(pbRows, &pb.MatrixRequest_Row{Values: row})
	}

	req := &pb.MatrixRequest{
		Rows: pbRows,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := c.client.ProcessRotatedMatrix(ctx, req)
	if err != nil {
		log.Printf("Error llamando a gRPC Node.js: %v", err)
		return nil, err
	}

	return res, nil
}
