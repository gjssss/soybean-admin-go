package utils

type Response[T any] struct {
	Code string `json:"code" binding:"required"`
	Data T      `json:"data" binding:"required"`
	Msg  string `json:"msg"  binding:"required"`
}

func NewResponse[T any](code string, data T, msg string) Response[T] {
	return Response[T]{Code: code, Data: data, Msg: msg}
}

func NewSuccessResponse[T any](data T) Response[T] {
	return NewResponse("0000", data, "success")
}
func NewErrorResponse(msg string) Response[string] {
	return NewResponse("0000", "", msg)
}
func NewLogoutModelResponse(code string, msg string) Response[string] {
	return NewResponse("8888", "", msg)
}
