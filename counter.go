package workers

import "sync"

type (
	counter struct {
		cnt int
		mtx sync.RWMutex
	}
)

func newCounter() *counter {
	return &counter{}
}

func (c *counter) apply(fn func(int) int) {
	c.mtx.Lock()
	c.cnt = fn(c.cnt)
	c.mtx.Unlock()
}

func (c *counter) incr(val int) {
	c.apply(func(i int) int {
		return i + val
	})
}

func (c *counter) get() int {
	c.mtx.RLock()
	val := c.cnt
	c.mtx.RUnlock()
	return val
}
