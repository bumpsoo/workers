package main

import (
	"context"
)

func WorkersWithFunc[ReqT any, ResT any](
	fn func(Request[ReqT]) Response[ResT], size int,
) Workers[ReqT, ResT] {
	if size <= 0 {
		size = 1
	}
	end := make(chan bool)
	workers := workers[ReqT, ResT]{
		workChan: make(chan coupled[ReqT, ResT], size),
		size:     size,
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
			case <-end:
				close(end)
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
