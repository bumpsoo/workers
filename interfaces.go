// Package workers contains to use funtions to worker pool easily with generic
// interfaces and structs
// You could wrap your functions in a anonymous function and pass it
package workers

import (
	"context"
)

type (
	// Use workers.WorkersWithFunc
	Workers[ReqT any, ResT any] interface {
		// Execute passed requests
		// Return closed channel
		Execute(...Request[ReqT]) []<-chan Response[ResT]

		// Max concurrency
		Size() int

		// Close
		Close()
	}
)

// Request interface
type Request[T any] interface {
	// Context
	Context() context.Context

	// Request
	Body() T
}

// Response interface
type Response[T any] interface {
	// Status
	Error() error

	// Response
	Body() T
}

// You can wrap ReqT and ResT with your interface and use it widely
type Manager[ReqT any, ResT any] interface {
	// Get workers
	Get(key string) (Workers[ReqT, ResT], error)

	// Put workers
	Put(key string, workers Workers[ReqT, ResT]) error

	// Close workers
	Close(key string) error

	// Execute request
	Execute(
		key string, request ...Request[ReqT],
	) ([]<-chan Response[ResT], error)
}
