package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/ujjwalmishra21/building-microservices-grpc/currency/protos/currency"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context, req *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", req.GetBase(), "destination", req.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
