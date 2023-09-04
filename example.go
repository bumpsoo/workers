package main

import (
	"context"
	"fmt"
	"time"
)

type Ex struct {
	Num int
	Str string
}

func main() {
	workers := WorkersWithFunc(
		func(r Request[Ex]) Response[string] {
			// doing some jobs.
			fmt.Println(r.Body().Str)
			return ResponseWithStatus(nil, r.Body().Str)
		},
		4,
	)
	workers.Execute(
		RequestWithCtx(Ex{Str: "foo"}, context.Background()),
		RequestWithCtx(Ex{Str: "foo"}, context.Background()),
		RequestWithCtx(Ex{Str: "foo"}, context.Background()),
		RequestWithCtx(Ex{Str: "foo"}, context.Background()),
	)
	time.Sleep(3 * time.Second)

}
