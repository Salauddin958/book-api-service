package repsitory

import (
	"context"

	"github.com/Salauddin958/book-api-service/models"
)

// Book Repository explain...
type BookRepo interface {
	Fetch(ctx context.Context, num int64) ([]*models.Book, error)
	GetByID(ctx context.Context, id int64) (*models.Book, error)
	Create(ctx context.Context, b *models.Book) (int64, error)
	Update(ctx context.Context, b *models.Book) (*models.Book, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
