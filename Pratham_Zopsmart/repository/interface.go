package repository

import (
  "gofr.dev/pkg/gofr"
  "Pratham_Zopsmart/model"
)

type BookRepository interface {

  CreateBook(ctx *gofr.Context, book *models.Book) error
  GetBook(ctx *gofr.Context, id int64) (*models.Book, error)
  UpdateBook(ctx *gofr.Context, book *models.Book) error
  DeleteBook(ctx *gofr.Context, id int64) error
  GetBooks(ctx *gofr.Context) ([]*models.Book, error) // added GetBooks()
}
