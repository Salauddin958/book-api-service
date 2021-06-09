package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Salauddin958/book-api-service/driver"
	models "github.com/Salauddin958/book-api-service/models"
	repository "github.com/Salauddin958/book-api-service/repository"
	book "github.com/Salauddin958/book-api-service/repository/book"
	"github.com/go-chi/chi"
)

// NewBookHandler ...
func NewBookHandler(db *driver.DB) *Book {
	return &Book{
		repo: book.NewSQLBookRepo(db.SQL),
	}
}

// Book ...
type Book struct {
	repo repository.BookRepo
}

// Fetch all Book data
func (b *Book) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, err := b.repo.Fetch(r.Context(), 5)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondwithJSON(w, http.StatusOK, payload)
}

// Create a new book
func (b *Book) Create(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	json.NewDecoder(r.Body).Decode(&book)

	newID, err := b.repo.Create(r.Context(), &book)
	fmt.Println(newID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Update a Book by id
func (b *Book) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.Book{ID: int(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := b.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a book details
func (b *Book) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := b.repo.GetByID(r.Context(), int64(id))
	if err != nil {
		respondWithError(w, http.StatusNoContent, err.Error())
		return
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a Book
func (b *Book) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := b.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
