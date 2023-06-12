package services

import (
	"log"

	"github.com/BryanAchmad/go-rest-api/internal/models"
	"github.com/BryanAchmad/go-rest-api/internal/repositories"
)

type Service interface {
	FindAll() ([]models.Book, error)
	FindOne(ID int) (models.Book, error)
	Create(book models.BookInput) (models.Book, error)
	Update(ID int, book models.BookInput) (models.Book, error)
}

type bookService struct{
	repository repositories.Repository
}

func NewBookService(repository repositories.Repository) Service {
	return &bookService{repository}
}

func (s *bookService) FindAll()([]models.Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *bookService) FindOne(ID int)(models.Book, error) {
	book, err := s.repository.FindOne(ID)
	log.Println(book)
	return book, err
}

func (s *bookService) Create(bookInput models.BookInput)(models.Book, error) {
	price, _ := bookInput.Price.Int64()

	book := models.Book{
		Title: bookInput.Title,
		Price: int(price),
		Description: bookInput.Description,
		Rating: bookInput.Rating,
		Discount: bookInput.Discount,
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *bookService) Update(ID int,bookInput models.BookInput)(models.Book, error) {
	book, err := s.repository.FindOne(ID)

	price, _ := bookInput.Price.Int64()

	book.Title = bookInput.Title
	book.Description = bookInput.Description
	book.Price = int(price)
	book.Rating = bookInput.Rating
	book.Discount =  bookInput.Discount

	newBook, err := s.repository.Update(book)
	return newBook, err
}