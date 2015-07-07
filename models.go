package main

// Book struct for the books table
type Book struct {
	ISBN   string  `gorm:"column:isbn;primary_key" json:"isbn"`
	Title  string  `gorm:"column:title" json:"title"`
	Author string  `gorm:"column:author" json:"author"`
	Price  float32 `gorm:"column:price" json:"price"`
}
