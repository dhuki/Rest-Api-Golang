package entity

// Book entity
type Book struct {
	// gorm.Model // basically it's just embedded struct / inheritance in go
	ID     string `json:"id" gorm:"primary_key;unique;not null"` // use tags to change name of key in format json
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}
