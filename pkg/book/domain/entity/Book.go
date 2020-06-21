package entity

// Book entity
type Book struct {
	// gorm.Model // basically it's just embedded struct / inheritance in go
	ID     string `gorm:"primary_key;unique;not null"` // use tags to change name of key in format json
	Title  string
	Author string
	Year   string
}
