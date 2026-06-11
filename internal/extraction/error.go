package extraction

type AppError struct {
	Kind    string
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NewError(kind, msg string) error {
	return AppError{Kind: kind, Message: msg}
}
