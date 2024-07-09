package entity

type UnauthorizedError struct {
    HttpStatus int
}

func (e UnauthorizedError) Error() string {
    return "user is unauthorized"
}

func (e UnauthorizedError) GetHttpStatus() int {
    return e.HttpStatus
}
