package counter_test

import (
	"testing"

	"github.com/bumpsoo/workers/counter"
)

func eqauls(val int, target int, t *testing.T) {
	if val != target {
		t.Error()
	}
}

func TestCounter(t *testing.T) {
	c := counter.NewCounter()
	val := c.Incr(3)
	eqauls(val, 3, t)
	val = c.Incr(-3)
	eqauls(val, 0, t)
	val = c.Get()
	eqauls(val, 0, t)
}
