package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kkato/book-api/models"
)

type BookHandler struct {
	store *models.BookStore
}

func NewBookHandler(store *models.BookStore) *BookHandler {
	return &BookHandler{store: store}
}

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listBooks(w, r)
	case http.MethodPost:
		h.createBook(w, r)
	default:
		h.respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	h.listBooks(w, r)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	h.createBook(w, r)
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id, ok := h.requireID(w, r)
	if !ok {
		return
	}
	h.getBook(w, r, id)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, ok := h.requireID(w, r)
	if !ok {
		return
	}
	h.updateBook(w, r, id)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, ok := h.requireID(w, r)
	if !ok {
		return
	}
	h.deleteBook(w, r, id)
}

func (h *BookHandler) HandleBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		h.respondError(w, http.StatusBadRequest, "book ID is required")
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getBook(w, r, id)
	case http.MethodPut:
		h.updateBook(w, r, id)
	case http.MethodDelete:
		h.deleteBook(w, r, id)
	default:
		h.respondError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (h *BookHandler) listBooks(w http.ResponseWriter, r *http.Request) {
	books := h.store.GetAll()
	h.respondJSON(w, http.StatusOK, books)
}

func (h *BookHandler) createBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := decodeJSONBody(r, &book); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if book.ID == "" {
		h.respondError(w, http.StatusBadRequest, "book ID is required")
		return
	}

	if err := h.store.Create(&book); err != nil {
		h.handleStoreError(w, err)
		return
	}

	h.respondJSON(w, http.StatusCreated, book)
}

func (h *BookHandler) getBook(w http.ResponseWriter, r *http.Request, id string) {
	book, err := h.store.GetByID(id)
	if err != nil {
		h.handleStoreError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, book)
}

func (h *BookHandler) updateBook(w http.ResponseWriter, r *http.Request, id string) {
	var book models.Book
	if err := decodeJSONBody(r, &book); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.store.Update(id, &book); err != nil {
		h.handleStoreError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, book)
}

func (h *BookHandler) deleteBook(w http.ResponseWriter, r *http.Request, id string) {
	if err := h.store.Delete(id); err != nil {
		h.handleStoreError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *BookHandler) requireID(w http.ResponseWriter, r *http.Request) (string, bool) {
	id := r.PathValue("id")
	if id == "" {
		h.respondError(w, http.StatusBadRequest, "book ID is required")
		return "", false
	}
	return id, true
}

func (h *BookHandler) respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func (h *BookHandler) respondError(w http.ResponseWriter, status int, message string) {
	h.respondJSON(w, status, map[string]string{"error": message})
}

func (h *BookHandler) handleStoreError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrBookNotFound):
		h.respondError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, models.ErrBookExists):
		h.respondError(w, http.StatusConflict, err.Error())
	default:
		h.respondError(w, http.StatusInternalServerError, "internal server error")
	}
}

func decodeJSONBody(r *http.Request, dest any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(dest); err != nil {
		return err
	}
	return nil
}
