package post

import (
	"context"
	"database/sql"

	models "github.com/Salauddin958/book-api-service/models"
	repo "github.com/Salauddin958/book-api-service/repository"
)

// NewSQLBookRepo returns implement of post repository interface
func NewSQLBookRepo(Conn *sql.DB) repo.BookRepo {
	return &mysqlBookRepo{
		Conn: Conn,
	}
}

type mysqlBookRepo struct {
	Conn *sql.DB
}

func (m *mysqlBookRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Book, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Book, 0)
	for rows.Next() {
		data := new(models.Book)

		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.Price,
			&data.Author,
			&data.Category,
			&data.Description,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlBookRepo) Fetch(ctx context.Context, num int64) ([]*models.Book, error) {
	query := "Select id, name, price, author, category, description From books limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlBookRepo) GetByID(ctx context.Context, id int64) (*models.Book, error) {
	query := "Select id, name, price, author, category, description From books where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Book{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlBookRepo) Create(ctx context.Context, b *models.Book) (int64, error) {
	query := "INSERT INTO books(name, price, author, category, description) VALUES(?, ?, ?, ?, ?)"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, b.Name, b.Price, b.Author, b.Category, b.Description)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlBookRepo) Update(ctx context.Context, b *models.Book) (*models.Book, error) {
	query := "UPDATE books SET name=?, price=?, author=?, category=?, description=? WHERE id=?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		b.Name,
		b.Price,
		b.Author,
		b.Category,
		b.Description,
		b.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return b, nil
}

func (m *mysqlBookRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From books Where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
