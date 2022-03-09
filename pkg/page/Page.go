package page

import "time"

type Page struct {
	PageSize   int         `json:"pageSize"`
	PageNumber int         `json:"pageNumber"`
	Sum        int         `json:"sum"`
	Data       interface{} `json:"data"`
	Timestamp  time.Time   `json:"timestamp"`
}

func (page *Page) OK(pageSize, pageNumber, sumPage int, data interface{}) Page {
	page.PageSize = pageSize
	page.PageNumber = pageNumber
	page.Sum = sumPage
	page.Data = data
	return *page
}
