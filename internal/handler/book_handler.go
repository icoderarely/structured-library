package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/icoderarely/structured-library/internal/domain"
	"github.com/icoderarely/structured-library/internal/service"
)

type BookHandler struct {
	bookService *service.BookService
}

func NewBookHandler(bs *service.BookService) *BookHandler {
	return &BookHandler{bookService: bs}
}

// GET    /books         → list all books
// GET    /books/{id}    → get a single book
// POST   /books         → add a new book
// DELETE /books/{id}    → delete a book

type envelope map[string]any

func writeJSON(w http.ResponseWriter, status int, data envelope) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, envelope{"error": message})
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book domain.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid req body")
		return
	}

	data, err := h.bookService.CreateBook(&book)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Error creating a book")
		return
	}
	writeJSON(w, http.StatusCreated, envelope{"book": data})
}

func (h *BookHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	book, err := h.bookService.GetBook(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}
	writeJSON(w, http.StatusOK, envelope{"book": book})
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.bookService.GetBooks()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error fetching books")
		return
	}
	writeJSON(w, http.StatusOK, envelope{"books": books})
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	err = h.bookService.DeleteBook(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "book not found")
		return
	}
	writeJSON(w, http.StatusOK, envelope{"message": "book deleted"})
}
