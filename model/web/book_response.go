package web

type BookResponse struct {
	Id            int    `json:"id"`
	CategoryId    int    `json:"category_id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	PublishedDate string `json:"published_date"`
	Price         int    `json:"price"`
	Stock         int    `json:"stock"`
}

type BookResponses struct {
	Id            int    `json:"id"`
	Category      string `json:"category"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	PublishedDate string `json:"published_date"`
	Price         int    `json:"price"`
	Stock         int    `json:"stock"`
}
