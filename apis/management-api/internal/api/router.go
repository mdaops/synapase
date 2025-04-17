package api

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
)

type Handler[I any, O any] func(ctx context.Context, input *I) (*O, error)

type Route struct {
	Op huma.Operation
	Fn Handler[any, any]
}
