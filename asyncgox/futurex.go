package asyncgox

import (
	"context"
)

const (
	single = 1
)

var _ Future = (*future)(nil)

// Future async/await programming model
type Future interface {
	Await(ctxs ...context.Context) (any, error)
}

type future struct {
	await AwaitFunc
}

// Await sync, await
func (f future) Await(ctxs ...context.Context) (any, error) {
	ctx := context.Background() // default: ctx
	switch len(ctxs) {
	case single:
		ctx = ctxs[0]
	}

	return f.await(ctx)
}
