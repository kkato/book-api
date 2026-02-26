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

func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	books := h.store.GetAll()
	h.respondJSON(w, http.StatusOK, books)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
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

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	book, err := h.store.GetByID(id)
	if err != nil {
		h.handleStoreError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, book)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

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

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.store.Delete(id); err != nil {
		h.handleStoreError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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
	return dec.Decode(dest)
}
