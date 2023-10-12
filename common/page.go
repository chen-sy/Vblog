package common

import (
	"strconv"

	"gorm.io/gorm"
)

// 默认第一页，每页10条数据
func NewPagination() *Pagination {
	return &Pagination{
		PageIndex: 1,
		PageSize:  10,
	}
}

// 分页数据
type Pagination struct {
	// 页码
	PageIndex int `json:"pageIndex"`
	// 每页大小
	PageSize int `json:"pageSize"`
}

// Scope 分页封装
func Paginate(pageIndex int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageIndex <= 0 {
			pageIndex = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (pageIndex - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (i *Pagination) ParsePageIndex(s string) {
	psInt, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i.PageIndex = int(psInt)
	}
}

func (i *Pagination) ParsePageSize(s string) {
	psInt, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i.PageSize = int(psInt)
	}
}
