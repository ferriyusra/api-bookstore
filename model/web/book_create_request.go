package web

type BookCreateRequest struct {
	CategoryId    int    `validate:"required,min=1,max=100" json:"category_id"`
	Title         string `validate:"required,min=1,max=100" json:"title"`
	Author        string `validate:"required,min=1,max=100" json:"author"`
	Publisher     string `validate:"required,min=1,max=100" json:"publisher"`
	PublishedDate string `validate:"required" json:"published_date"`
	Price         int    `validate:"required" json:"price"`
	Stock         int    `validate:"required" json:"stock"`
}
