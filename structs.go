package workers

import (
	"context"
	"sync"
)

type (
	coupled[ReqT any, ResT any] struct {
		request      Request[ReqT]
		responseChan chan Response[ResT]
	}

	workers[ReqT any, ResT any] struct {
		workChan chan coupled[ReqT, ResT]
		size     int
		end      chan bool
		mtx      sync.Mutex
		cnt      int
	}

	request[T any] struct {
		body T
		ctx  context.Context
	}

	response[T any] struct {
		body T
		err  error
	}

	manager[ReqT any, ResT any] struct {
		workers map[string]Workers[ReqT, ResT]
	}
)

func (workers workers[req, res]) Size() int {
	return workers.size
}

func (worker workers[ReqT, ResT]) Execute(
	request ...Request[ReqT],
) []<-chan Response[ResT] {
	ret := make([]<-chan Response[ResT], len(request))
	for _, value := range request {
		channel := make(chan Response[ResT], 1)
		worker.workChan <- coupled[ReqT, ResT]{
			request:      value,
			responseChan: channel,
		}
		ret = append(ret, channel)
	}
	return ret
}

func (worker workers[req, res]) Close() {
	close(worker.end)
}

func (request request[T]) Context() context.Context {
	return request.ctx
}

func (request request[T]) Body() T {
	return request.body
}

func (man manager[ReqT, ResT]) Put(
	key string, worker Workers[ReqT, ResT],
) error {
	_, err := man.Get(key)
	if err != nil {
		return newError(MANAGER_KEY_OCCUPIED)
	}
	man.workers[key] = worker
	return nil
}

func (man manager[ReqT, ResT]) Get(key string) (Workers[ReqT, ResT], error) {
	worker, ok := man.workers[key]
	if !ok {
		return nil, newError(MANAGER_KEY_NOT_FOUND)
	}
	return worker, nil
}

func (man manager[ReqT, ResT]) Execute(
	key string, request ...Request[ReqT],
) ([]<-chan Response[ResT], error) {
	worker, err := man.Get(key)
	if err != nil {
		return nil, err
	}

	return worker.Execute(request...), nil
}
func (man manager[req, res]) Close(key string) error {
	worker, err := man.Get(key)
	if err == nil {
		worker.Close()
		delete(man.workers, key)
	}
	return err
}

func (man manager[req, res]) Count() int {
	return len(man.workers)
}

func (response response[T]) Error() error {
	return response.err
}

func (response response[T]) Body() T {
	return response.body
}
