package workers

type (
	workers[ReqT any, ResT any] struct {
		fn     work[ReqT, ResT]
		size   int
		cnt    *counter
		closed bool
	}
)

func (w workers[req, res]) Size() int {
	return w.size
}

func (w workers[ReqT, ResT]) Count() int {
	return w.cnt.cnt
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
			w.cnt.incr(1)
			res := w.fn(req)
			responseChan <- res
			close(responseChan)
			w.cnt.incr(-1)
		}(value, channel)
		ret = append(ret, channel)
	}
	return ret
}

func (w *workers[req, res]) Close() {
	cnt := w.cnt.get()
	if cnt > 0 {
		w.Close()
	} else {
		w.closed = true
	}
}
