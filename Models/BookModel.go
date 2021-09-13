package Models

// book represents a book.
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func (b *Book) TableName() string {
	return "book"
}
