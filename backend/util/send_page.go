package util

import (
	"net/http"
)

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"total_items"`
	TotalPages int64 `json:"total_pages"`
}

func SendPage(w http.ResponseWriter, data any, page, limit, cnt int64) {

	//input validation
	if page < 1 {
		page = 1
	}
	// by default 10 items per page
	if limit <= 0 {
		limit = 10
	}

	//page calculation
	totalPages := (cnt + limit - 1) / limit
	if totalPages == 0 {
		totalPages = 1
	}

	// page should not exceed total pages
	if page > totalPages {
		page = totalPages
	}

	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: cnt,
			TotalPages: totalPages,
		},
	}

	SendData(w, http.StatusOK, paginatedData)
}
