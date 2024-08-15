package server

import (
	"context"

	"github.com/hashicorp/go-hclog"

	protos "github.com/varun-muthanna/products-api/currency"
)

type Currency struct {
	log hclog.Logger
	protos.UnimplementedCurrencyServer
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("handle getrate", "base", rr.GetSource(), "destination", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}
