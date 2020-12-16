package services

import (
	"context"
	"fmt"
)

// ProdService empty struct should implememt `ProdServiceServer` and `ProdServiceClient` interface.
type ProdService struct {
}

// GetProdStock implement the protobuf file
// `ProdServiceServer` and `ProdServiceClient` interface.
func (thisSvc *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {
	switch request.ProdId {
	case 12:
			return &ProdResponse{ProdStock: 20, ProdName: "Pencil"}, nil
	case 13:
			return &ProdResponse{ProdStock: 10, ProdName:"Pencil Box"}, nil
	default:
		  return &ProdResponse{}, fmt.Errorf("Unknown ProdId %d", request.ProdId)
		}
}
