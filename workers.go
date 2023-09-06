package workers

import "github.com/bumpsoo/workers/counter"

type (
	workers[ReqT any, ResT any] struct {
		fn      work[ReqT, ResT]
		size    int
		counter *counter.Counter
		closed  bool
	}
)

func StartWorkers[ReqT any, ResT any](
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

func (w workers[req, res]) Size() int {
	return w.size
}

func (w workers[ReqT, ResT]) Count() int {
	return w.counter.Get()
}

func (w workers[ReqT, ResT]) IsClosed() bool {
	return w.closed
}

func (w workers[ReqT, ResT]) Execute(
	request []Request[ReqT],
) []<-chan Response[ResT] {
	if !w.closed {
		return nil
	}
	ret := make([]<-chan Response[ResT], len(request))
	for _, value := range request {
		channel := make(chan Response[ResT], 1)
		go func(req Request[ReqT], responseChan chan Response[ResT]) {
			w.counter.Incr(1)
			res := w.fn(req)
			responseChan <- res
			close(responseChan)
			w.counter.Incr(-1)
		}(value, channel)
		ret = append(ret, channel)
	}
	return ret
}

func (w *workers[req, res]) Close() {
	cnt := w.counter.Get()
	if cnt > 0 {
		w.Close()
	} else {
		w.closed = true
	}
}
