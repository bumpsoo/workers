package workers

import (
	"sync"

	"github.com/bumpsoo/workers/counter"
)

type (
	manager[ReqT any, ResT any] struct {
		pool sync.Map
		cnt  *counter.Counter
	}
)

func StartManager[ReqT any, ResT any]() Manager[ReqT, ResT] {
	return &manager[ReqT, ResT]{
		pool: sync.Map{},
	}
}

func (m manager[ReqT, ResT]) Put(
	key any, worker Workers[ReqT, ResT],
) error {
	_, loaded := m.pool.LoadOrStore(key, worker)
	if loaded {
		return newError(MANAGER_KEY_OCCUPIED)
	} else {
		m.cnt.Incr(1)
	}
	return nil
}

func (m manager[ReqT, ResT]) Get(key any) (Workers[ReqT, ResT], error) {
	ret, ok := m.pool.Load(key)
	if !ok {
		return nil, newError(MANAGER_KEY_NOT_FOUND)
	}
	return ret.(Workers[ReqT, ResT]), nil
}

func (m manager[ReqT, ResT]) Execute(
	key any, request []Request[ReqT],
) ([]<-chan Response[ResT], error) {
	workers, err := m.Get(key)
	if err != nil {
		return nil, err
	}
	return workers.Execute(request), nil
}

func (m manager[ReqT, ResT]) Close(key any) error {
	workers, err := m.Get(key)
	if err == nil {
		workers.Close()
		m.pool.Delete(key)
	}
	return err
}

func (m manager[ReqT, ResT]) Count() int {
	val := m.cnt.Get()
	return val
}
