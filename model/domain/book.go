package domain

type BookCreateOrUpdate struct {
	Id            int
	CategoryId    int
	Title         string
	Author        string
	Publisher     string
	PublishedDate string
	Price         int
	Stock         int
}
type Book struct {
	Id            int
	CategoryId    int
	Category      string
	Title         string
	Author        string
	Publisher     string
	PublishedDate string
	Price         int
	Stock         int
}
