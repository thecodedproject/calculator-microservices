package add

import (
	"context"
)

type Client interface {
	Calc(ctx context.Context, values []float64) (float64, error)
}
