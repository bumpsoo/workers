package workers

import (
	"context"

	"github.com/bumpsoo/workers/counter"
)

type Work[ReqT any, ResT any] interface {
	Start(size int) Workers[ReqT, ResT]

	execute(request Request[ReqT]) Response[ResT]
}

type work[ReqT any, ResT any] func(Request[ReqT]) Response[ResT]

func NewWork[ReqT any, ResT any](
	fn func(ctx context.Context, req ReqT) ResT,
) Work[ReqT, ResT] {
	return work[ReqT, ResT](func(r Request[ReqT]) Response[ResT] {
		res := fn(r.Context(), r.Body())
		return InitResponse(res, nil)
	})
}

func NewWorkWithErr[ReqT any, ResT any](
	fn func(ctx context.Context, req ReqT) (ResT, error),
) Work[ReqT, ResT] {
	return work[ReqT, ResT](
		func(r Request[ReqT]) Response[ResT] {
			res, err := fn(r.Context(), r.Body())
			return InitResponse(res, err)
		})
}

func (w work[ReqT, ResT]) Start(size int) Workers[ReqT, ResT] {
	if size <= 0 {
		size = 1
	}
	workers := &workers[ReqT, ResT]{
		fn:      w,
		size:    size,
		counter: counter.NewCounter(),
		closed:  false,
	}
	return workers
}

func (w work[ReqT, ResT]) execute(request Request[ReqT]) Response[ResT] {
	return w(request)
}
