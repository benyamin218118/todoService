package domain

type BadRequestError struct {
	Msg string
}

func (e BadRequestError) Error() string {
	return e.Msg
}

type ForbiddenError struct {
	Msg string
}

func (e ForbiddenError) Error() string {
	return e.Msg
}
