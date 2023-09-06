package workers

import (
	"context"
	"sync"

	"github.com/bumpsoo/workers/counter"
)

type work[ReqT any, ResT any] func(Request[ReqT]) Response[ResT]

func WorkersWithFunc[ReqT any, ResT any](
	fn work[ReqT, ResT], size int,
) Workers[ReqT, ResT] {
	if size <= 0 {
		size = 1
	}
	workers := &workers[ReqT, ResT]{
		fn:      fn,
		size:    size,
		counter: counter.NewCounter(),
	}
	return workers
}

func ResponseWithError[T any](err error, body T) Response[T] {
	return &response[T]{
		body: body,
		err:  err,
	}
}

func StartManager[ReqT any, ResT any]() Manager[ReqT, ResT] {
	return &manager[ReqT, ResT]{
		pool: sync.Map{},
	}
}

func RequestWithCtx[ReqT any](
	body ReqT, ctx context.Context,
) Request[ReqT] {
	return &request[ReqT]{
		body: body,
		ctx:  ctx,
	}
}
