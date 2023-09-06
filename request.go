package workers

import "context"

type (
	request[T any] struct {
		body T
		ctx  context.Context
	}
)

func DefaultRequest[T any](body T) Request[T] {
	return RequestWithCtx(body, context.Background())
}

func RequestWithCtx[T any](body T, ctx context.Context) Request[T] {
	return &request[T]{body: body, ctx: ctx}
}

func (r *request[T]) Context() context.Context {
	return r.ctx
}

func (r *request[T]) Body() T {
	return r.body
}
