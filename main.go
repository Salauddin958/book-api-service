package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Salauddin958/book-api-service/driver"
	ph "github.com/Salauddin958/book-api-service/handler/http"
	"github.com/go-chi/chi"
)

func main() {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// dbName := "booksdb"
	// dbPass := "root"
	// dbHost := "localhost"
	//dbPort := "3306"

	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()

	pHandler := ph.NewBookHandler(connection)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/books", bookRouter(pHandler))
	})

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":8005", r)
}

// A completely separate router for books routes
func bookRouter(pHandler *ph.Book) http.Handler {
	r := chi.NewRouter()
	r.Get("/", pHandler.Fetch)
	r.Get("/{id:[0-9]+}", pHandler.GetByID)
	r.Post("/", pHandler.Create)
	r.Put("/{id:[0-9]+}", pHandler.Update)
	r.Delete("/{id:[0-9]+}", pHandler.Delete)

	return r
}
