package workers

type (
	response[T any] struct {
		body T
		err  error
	}
)

func (r *response[T]) Error() error {
	return r.err
}

func (r *response[T]) Body() T {
	return r.body
}
