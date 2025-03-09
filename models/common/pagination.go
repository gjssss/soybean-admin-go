package common

type PaginationParam struct {
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}

type Pagination[T any] struct {
	Records []T `json:"records"`
	Current int `json:"current"`
	Size    int `json:"size"`
	Total   int `json:"total"`
}
