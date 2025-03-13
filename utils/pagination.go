package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaginationParam struct {
	Current int `json:"current"`
	Size    int `json:"size"`
}

type Pagination[T any] struct {
	Records []T   `json:"records" binding:"required"`
	Current int   `json:"current" binding:"required"`
	Size    int   `json:"size"    binding:"required"`
	Total   int64 `json:"total"   binding:"required"`
}

func ParsePagination(ctx *gin.Context) PaginationParam {
	Current, _ := strconv.Atoi(ctx.DefaultQuery("current", "1"))
	Size, _ := strconv.Atoi(ctx.DefaultQuery("size", "10"))
	return PaginationParam{Current: Current, Size: Size}
}

func NewPagination[T any](data []T, page PaginationParam, total int64) Pagination[T] {
	return Pagination[T]{
		Current: page.Current,
		Size:    page.Size,
		Total:   total,
		Records: data,
	}
}
