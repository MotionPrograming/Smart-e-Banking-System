package util

import (
	"net/http"
)

type Pagination struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"total_items"`
	TotalPages int64 `json:"total_pages"`
}

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func SendPage(w http.ResponseWriter, data any, page, limit, cnt int64) {
	// ইনপুট ভ্যালিডেশন ও ডিফল্ট ভ্যালু সেটআপ
	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	// মোট পেজ সংখ্যা হিসাব করার নিখুঁত সূত্র
	totalPages := (cnt + limit - 1) / limit
	if totalPages == 0 {
		totalPages = 1
	}

	// রিকোয়েস্ট করা পেজ যদি টোটাল পেজের চেয়ে বড় হয়, তবে শেষ পেজে রিডাইরেক্ট করা
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
