package workers

type (
	response[T any] struct {
		body T
		err  error
	}
)

func Success[T any](body T) Response[T] {
	return responseWithError(body, nil)
}

func Fail[T any](body T, err error) Response[T] {
	return responseWithError(body, err)
}

func responseWithError[T any](body T, err error) Response[T] {
	return &response[T]{body: body, err: err}
}

func (r *response[T]) Error() error {
	return r.err
}

func (r *response[T]) Body() T {
	return r.body
}
