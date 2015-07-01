package main

// Book struct for the books table
type Book struct {
	ISBN   string  `gorm:"column:isbn;primary_key"`
	Title  string  `gorm:"column:title"`
	Author string  `gorm:"column:author"`
	Price  float32 `gorm:"column:price"`
}
