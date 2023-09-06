package counter

import "sync"

type (
	Counter struct {
		cnt int
		mtx sync.RWMutex
	}
)

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) apply(fn func(int) int) int {
	c.mtx.Lock()
	val := fn(c.cnt)
	c.cnt = val
	c.mtx.Unlock()
	return val
}

func (c *Counter) Incr(val int) int {
	return c.apply(func(i int) int {
		return i + val
	})
}

func (c *Counter) Get() int {
	c.mtx.RLock()
	val := c.cnt
	c.mtx.RUnlock()
	return val
}
