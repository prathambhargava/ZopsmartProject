package handler

import (
	"encoding/json"
	"strconv"

	"gofr.dev/pkg/gofr"

	"Pratham_Zopsmart/model"
	"Pratham_Zopsmart/repository"
)

type BookHandler struct {
	repo repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{repo: repo}
}

func (h *BookHandler) GetBooks(ctx *gofr.Context) (interface{}, error) {
	books, err := h.repo.GetBooks(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (h *BookHandler) CreateBook(ctx *gofr.Context) (interface{}, error) {
	var book models.Book
	err := json.NewDecoder(ctx.Request().Body).Decode(&book)
	if err != nil {
		return nil, err
	}

	err = h.repo.CreateBook(ctx, &book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (h *BookHandler) GetBook(ctx *gofr.Context) (interface{}, error) {
	id, err := strconv.ParseInt(ctx.PathParam("id"), 10, 64)
	if err != nil {
		return nil, err
	}

	book, err := h.repo.GetBook(ctx, id)
	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil
	}

	return book, nil
}

func (h *BookHandler) UpdateBook(ctx *gofr.Context) (interface{}, error) {
	id, err := strconv.ParseInt(ctx.PathParam("id"), 10, 64)
	if err != nil {
		return nil, err
	}

	var book models.Book
	err = json.NewDecoder(ctx.Request().Body).Decode(&book)
	if err != nil {
		return nil, err
	}

	book.ID = id
	err = h.repo.UpdateBook(ctx, &book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (h *BookHandler) DeleteBook(ctx *gofr.Context) (interface{}, error) {
	id, err := strconv.ParseInt(ctx.PathParam("id"), 10, 64)
	if err != nil {
		return nil, err
	}

	err = h.repo.DeleteBook(ctx, id)
	if err != nil {
		return nil, err
	}

	return nil
}
