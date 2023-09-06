// Package workers contains to use funtions to worker pool easily with generic
// interfaces and structs
// You could wrap your functions in a anonymous function and pass it
package workers

import (
	"context"
)

type (
	work[ReqT any, ResT any] func(Request[ReqT]) Response[ResT]
	// Use workers.WorkersWithFunc
	Workers[ReqT any, ResT any] interface {
		// Execute passed requests
		// Return closed channel
		Execute([]Request[ReqT]) []<-chan Response[ResT]

		// Max concurrency
		Size() int

		// Close
		Close()

		// check whether workers closed
		IsClosed() bool

		// get how many workers are working
		Count() int
	}

	// Request interface
	Request[T any] interface {
		// Context
		Context() context.Context

		// Request
		Body() T
	}

	// Response interface
	Response[T any] interface {
		// Status
		Error() error

		// Response
		Body() T
	}

	// You can wrap ReqT and ResT with your interface and use it
	Manager[ReqT any, ResT any] interface {
		// Get Workers
		Get(key any) (Workers[ReqT, ResT], error)

		// Put Workers
		Put(key any, workers Workers[ReqT, ResT]) error

		// Close Workers
		Close(key any) error

		// Execute Requests
		Execute(key any, req []Request[ReqT]) ([]<-chan Response[ResT], error)

		// How many Workers has been put
		Count() int
	}
)
