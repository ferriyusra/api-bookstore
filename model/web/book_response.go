package web

type BookResponse struct {
	Id            int    `json:"id"`
	CategoryId    int    `json:"category_id"`
	CategoryName  string `json:"category_name"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	PublishedDate string `json:"published_date"`
	Price         int    `json:"price"`
	Stock         int    `json:"stock"`
}
