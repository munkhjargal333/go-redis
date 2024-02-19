package shared

import "github.com/gofiber/fiber/v2"

type Pagination[T any] struct {
	Offset     int   `json:"-"`
	PageSize   int   `json:"page_size"`
	PageNumber int   `json:"page_number"`
	TotalPage  int64 `json:"total_page"`
	TotalRow   int64 `json:"total_row"`
	Items      []T   `json:"items"`
}

func Paginate[T any](c *fiber.Ctx, totalRows int64) *Pagination[T] {
	pageSize := c.QueryInt("page_size", 50)
	pageNumber := c.QueryInt("page_number", 1)

	offset := pageSize * (pageNumber - 1)
	pagination := &Pagination[T]{
		PageSize:   pageSize,
		PageNumber: pageNumber,
		TotalRow:   totalRows,
		Offset:     offset,
		TotalPage:  (totalRows + int64(pageSize) - 1) / int64(pageSize),
	}
	return pagination
}
