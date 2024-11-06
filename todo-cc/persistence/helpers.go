package persistence

type DbConnectionError struct {
	Message string
}

func (e *DbConnectionError) Error() string {
	return e.Message
}