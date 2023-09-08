package workers

type (
	response[T any] struct {
		body T
		err  error
	}
)

func InitResponse[T any](body T, err error) Response[T] {
	return &response[T]{body: body, err: err}
}

func (r *response[T]) Error() error {
	return r.err
}

func (r *response[T]) Body() T {
	return r.body
}
