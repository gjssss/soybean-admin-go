package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationParam struct {
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}

type Pagination[T any] struct {
	Records []T   `json:"records"`
	Current int   `json:"current"`
	Size    int   `json:"size"`
	Total   int64 `json:"total"`
}

func ParsePagination(ctx *gin.Context) PaginationParam {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	return PaginationParam{Current: page, PageSize: pageSize}
}

func NewPagination[T any](data []T, page PaginationParam, total int64) Pagination[T] {
	return Pagination[T]{
		Current: page.Current,
		Size:    page.PageSize,
		Total:   total,
		Records: data,
	}
}
