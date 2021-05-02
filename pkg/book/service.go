package book

import "go-crud-dynamodb-gofiber/pkg/entities"

type Service interface {
	InsertBook(book *entities.Book) (*entities.Book, error)
	FetchBooks() (*[]entities.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	RemoveBook(book *entities.Book) error
}

type service struct {
	repo Repository
}

func (s *service) InsertBook(book *entities.Book) (*entities.Book, error) {
	return s.repo.CreateBook(book)
}

func (s *service) FetchBooks() (*[]entities.Book, error) {
	return s.repo.ReadBooks()
}

func (s *service) UpdateBook(book *entities.Book) (*entities.Book, error) {
	return s.repo.UpdateBook(book)
}

func (s *service) RemoveBook(book *entities.Book) error {
	return s.repo.DeleteBook(book)
}

func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}
