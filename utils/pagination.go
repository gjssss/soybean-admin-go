package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gjssss/soybean-admin-go/models/common"
)

func ParsePagination(ctx *gin.Context) common.PaginationParam {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	return common.PaginationParam{Current: page, PageSize: pageSize}
}

func NewPagination[T any](data []T, page common.PaginationParam, total int64) common.Pagination[T] {
	return common.Pagination[T]{
		Current: page.Current,
		Size:    page.PageSize,
		Total:   total,
		Records: data,
	}
}
