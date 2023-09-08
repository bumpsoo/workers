package workers_test

import (
	"context"
	"testing"

	"github.com/bumpsoo/workers"
)

type my_key string

const db_worker = my_key("hello world")

func TestManager(t *testing.T) {
	m := workers.StartManager[int, string]()
	tt := temp{str: "who am i"}
	f := func(ctx context.Context, i int) string {
		ret := tt.foo(i)
		t.Log(ret)
		return ret
	}
	err := m.Put(db_worker, workers.NewWork(f).Start(3))
	if err != nil {
		t.Fatal(err.Error())
	}
	jobs := []workers.Request[int]{}
	for i := 0; i < 10; i++ {
		jobs = append(jobs, workers.DefaultRequest(i))
	}
	channel, err := m.Execute(db_worker, jobs)
	if channel == nil || err != nil {
		t.Fatal()
	}
	for _, ch := range channel {
		if ch == nil {
			t.Fatal("nil channel")
		}
		t.Log((<-ch).Body())
	}

}
