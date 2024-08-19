package client

import (
	"context"
	"errors"
	"fmt"

	"github.com/reuben-emmens/revisio/internal/models/csvfile"
)

var (
	ErrNewClient  = errors.New("unable to create a client")
	ErrClientType = errors.New("client type requested is invalid")
)

type Client interface {
	Creater
	Reader
}

type Creater interface {
	Create(ctx context.Context, key, value string) error
}

type Reader interface {
	ReadValue(ctx context.Context, key string) (string, error)
}

func New(ctx context.Context, clientType string) (Client, error) {
	switch clientType {
	case "csvfile":
		csvfilePtr, err := csvfile.New(ctx)
		if err != nil {
			fmt.Printf("%s", err)
		}
		return csvfilePtr, nil
	}
	return nil, ErrClientType
}
