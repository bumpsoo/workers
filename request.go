package workers

import "context"

type (
	request[T any] struct {
		body T
		ctx  context.Context
	}
)

func (r *request[T]) Context() context.Context {
	return r.ctx
}

func (r *request[T]) Body() T {
	return r.body
}
