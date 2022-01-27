package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
type CategoryWebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
type BookWebResponse struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	CountData int         `json:"count_data"`
	Data      interface{} `json:"data"`
}
