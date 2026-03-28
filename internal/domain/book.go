package domain

type Book struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Author      string `json:"author" db:"author"`
	Publisher   string `json:"publisher" db:"publisher"`
	PublishDate string `json:"publish_date" db:"publish_date"`
}
