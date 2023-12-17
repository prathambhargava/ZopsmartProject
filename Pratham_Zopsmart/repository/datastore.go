package repository

import (
  "database/sql"
  "errors"
  "fmt"

  _ "github.com/go-sql-driver/mysql" // MySQL driver
  "gofr.dev/pkg/gofr"
  "Pratham_Zopsmart/model"
)

type MySQLRepository struct {
  app *gofr.App
}

func NewMySQLRepository(app *gofr.App) (BookRepository, error) {
  db := app.Get("db").(*sql.DB) // access DB connection from Gofr App
  if db == nil {
    return nil, errors.New("missing database connection from Gofr app")
  }
  return &MySQLRepository{app: app, db: db}, nil
}

func (repo *MySQLRepository) CreateBook(ctx *gofr.Context, book *models.Book) error {
  // Prepare statement with named parameters
  stmt, err := repo.db.PrepareContext(ctx.Context(), "INSERT INTO books (name, author, student_id, status) VALUES (?, ?, ?, ?)")
  if err != nil {
    return err
  }
  defer stmt.Close()

  // Execute statement with book values
  _, err = stmt.ExecContext(ctx.Context(), book.Name, book.Author, book.StudentID, book.Status)
  if err != nil {
    return err
  }
  // Get newly generated ID if needed
  if book.ID == 0 {
    lastID, err := repo.db.LastInsertId(ctx.Context())
    if err != nil {
      return err
    }
    book.ID = lastID
  }

  return nil
}

func (repo *MySQLRepository) GetBook(ctx *gofr.Context, id int64) (*models.Book, error) {
  // Prepare statement with named parameter
  stmt, err := repo.db.PrepareContext(ctx.Context(), "SELECT * FROM books WHERE id = ?")
  if err != nil {
    return nil, err
  }
  defer stmt.Close()

  var book models.Book

  // Scan returned row into book struct
  err = stmt.QueryRowContext(ctx.Context(), id).Scan(&book.ID, &book.Name, &book.Author, &book.StudentID, &book.Status)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return nil, nil // not found
    }
    return nil, err
  }

  return &book, nil
}

func (repo *MySQLRepository) UpdateBook(ctx *gofr.Context, book *models.Book) error {
  // Prepare statement with named parameters
  stmt, err := repo.db.PrepareContext(ctx.Context(), "UPDATE books SET name = ?, author = ?, student_id = ?, status = ? WHERE id = ?")
  if err != nil {
    return err
  }
  defer stmt.Close()

  // Execute statement with updated book values and ID
  _, err = stmt.ExecContext(ctx.Context(), book.Name, book.Author, book.StudentID, book.Status, book.ID)
  if err != nil {
    return err
  }

  return nil
}

func (repo *MySQLRepository) DeleteBook(ctx *gofr.Context, id int64) error {
  // Prepare statement with named parameter
  stmt, err := repo.db.PrepareContext(ctx.Context(), "DELETE FROM books WHERE id = ?")
  if err != nil {
    return err
  }
  defer stmt.Close()

  // Execute statement with book ID
  _, err = stmt.ExecContext(ctx.Context(), id)
  if err != nil {
    return err
  }

  return nil
}

func (repo *MySQLRepository) GetBooks(ctx *gofr.Context) ([]*models.Book, error) {
  // Prepare statement for selecting all books
  stmt, err := repo.db.PrepareContext(ctx.Context(), "SELECT * FROM books")
  if err != nil {
    return nil, err
  }
  defer stmt.Close()

  rows, err := stmt.QueryContext(ctx.Context())
  if err != nil {
    return nil
  }