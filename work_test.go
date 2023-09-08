package workers_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/bumpsoo/workers"
)

type temp struct {
	str string
}

func (t temp) foo(i int) string {
	ret := strconv.Itoa(i) + t.str
	return ret
}

func TestWork(t *testing.T) {
	tt := temp{str: "hello world"}
	f := func(ctx context.Context, i int) string {
		ret := tt.foo(i)
		t.Log(ret)
		return ret
	}
	w := workers.NewWork[int, string](f).Start(3)
	jobs := []workers.Request[int]{}
	for i := 0; i < 10; i++ {
		jobs = append(jobs, workers.DefaultRequest(i))
	}
	channel := w.Execute(jobs)
	if channel == nil {
		t.Fatal("nil channel")
	}
	for _, ch := range channel {
		if ch == nil {
			t.Fatal("nil channel")
		}
		t.Log((<-ch).Body())
	}
}
