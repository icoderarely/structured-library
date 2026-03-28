package domain

type BookRepository interface {
	GetById(ID int) (*Book, error)
	Create(book *Book) (*Book, error)
	List() ([]*Book, error)
	Delete(ID int) error
}
