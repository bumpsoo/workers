package workers

import (
	"context"
	"sync"
)

type work[ReqT any, ResT any] func(Request[ReqT]) Response[ResT]

func WorkersWithFunc[ReqT any, ResT any](
	fn work[ReqT, ResT], size int,
) Workers[ReqT, ResT] {
	if size <= 0 {
		size = 1
	}
	workers := workers[ReqT, ResT]{
		workChan: make(chan coupled[ReqT, ResT], size),
		size:     size,
		end:      make(chan bool),
		mtx:      sync.Mutex{},
		cnt:      0,
	}
	go func() {
		for {
			select {
			case val := <-workers.workChan:
				go func() {
					res := fn(val.request)
					val.responseChan <- res
					close(val.responseChan)
				}()
			case <-workers.end:
				close(workers.workChan)
			}
		}
	}()
	return workers
}

func ResponseWithStatus[T any](err error, body T) Response[T] {
	return response[T]{
		body: body,
		err:  err,
	}
}

func StartManager[ReqT any, ResT any]() Manager[ReqT, ResT] {
	return manager[ReqT, ResT]{
		workers: map[string]Workers[ReqT, ResT]{},
	}
}

func RequestWithCtx[ReqT any](
	body ReqT, ctx context.Context,
) Request[ReqT] {
	return request[ReqT]{
		body: body,
		ctx:  ctx,
	}
}
