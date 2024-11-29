package shared

type DbConnectionError struct {
	Message string
}

type ExecError struct {
	Message string
}

func (e *DbConnectionError) Error() string {
	return e.Message
}

func (dbe *ExecError) Error() string {
	return dbe.Message
}
