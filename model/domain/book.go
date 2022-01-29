package domain

type Book struct {
	Id            int
	CategoryId    int
	CategoryName  string
	Title         string
	Author        string
	Publisher     string
	PublishedDate string
	Price         int
	Stock         int
}
