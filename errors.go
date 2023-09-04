package workers

type errorString string

type workerError struct {
	s errorString
}

func (e workerError) Error() string {
	return string(e.s)
}

func newError(str errorString) error {
	return workerError{str}
}
