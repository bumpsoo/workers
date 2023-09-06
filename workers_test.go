package workers_test

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bumpsoo/workers"
)

func TestWorkers(t *testing.T) {
	w := workers.StartWorkers(
		func(r workers.Request[int]) workers.Response[string] {
			str := "foo" + strconv.Itoa(r.Body())
			return workers.Success(str)
		},
		4,
	)
	reqs := []workers.Request[int]{}
	for i := 0; i < 8; i++ {
		reqs = append(reqs, workers.DefaultRequest(i))
	}
	ret := w.Execute(reqs)
	if ret == nil {
		t.Fatal("nil channel")
	}
	t.Log("length: ", len(ret))
	time.Sleep(1 * time.Second)
	for _, ch := range ret {
		if ch == nil {
			t.Fatal("nil channel")
		}
		res := <-ch
		if !strings.Contains(res.Body(), "foo") {
			t.Error("bar")
		}
	}

}
