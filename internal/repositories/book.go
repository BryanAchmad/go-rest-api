package repositories

import (
	"github.com/BryanAchmad/go-rest-api/internal/models"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Book, error)
	FindOne(ID int) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book) (models.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.Book, error) {
	var book []models.Book
	err := r.db.Find(&book).Error

	return book, err
}

func (r *repository) FindOne(ID int) (models.Book, error) {
	var book models.Book
	err:= r.db.Find(&book, ID).Error

	return book, err
}

func (r *repository) Create(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) Update(book models.Book) (models.Book, error) {
	err := r.db.Save(&book).Error

	return book, err
}