package http

type Error struct {
	err error
}

// Error
// 实现error interface
func (e Error) Error() string {
	return e.err.Error()
}
