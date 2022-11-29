package repository

import (
	"cake-api/core"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func NewCakeRepository(db *sql.DB) core.CakeRepository {
	return &cakeRepositoryImpl{db: db}
}

type cakeRepositoryImpl struct {
	db *sql.DB
}

func scanRowToCake(rows *sql.Rows, cake *core.Cake) error {
	return rows.Scan(
		&cake.ID,
		&cake.Title,
		&cake.Description,
		&cake.Rating,
		&cake.Image,
		&cake.CreatedAt,
		&cake.UpdatedAt)
}

func (r *cakeRepositoryImpl) FindCakeList(ctx context.Context) ([]core.Cake, error) {
	rows, err := r.db.Query("SELECT * FROM cakes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cakes []core.Cake
	for rows.Next() {
		cake := core.Cake{}
		if err := scanRowToCake(rows, &cake); err != nil {
			return nil, err
		}
		cakes = append(cakes, cake)
	}
	return cakes, nil
}

func (r *cakeRepositoryImpl) FindCakeByID(ctx context.Context, cakeID uint) (*core.Cake, error) {
	row := r.db.QueryRow("SELECT * FROM cakes WHERE id = ?", cakeID)
	if err := row.Err(); err != nil {
		return nil, err
	}

	cake := core.Cake{}
	err := row.Scan(
		&cake.ID,
		&cake.Title,
		&cake.Description,
		&cake.Rating,
		&cake.Image,
		&cake.CreatedAt,
		&cake.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &cake, nil
}

func (r *cakeRepositoryImpl) SaveCake(ctx context.Context, cake *core.Cake) (*core.Cake, error) {
	updatedAt := time.Now()
	createdAt := cake.CreatedAt
	if createdAt.IsZero() {
		createdAt = updatedAt
	}

	result, err := r.db.Exec(`
			INSERT INTO cakes(id, title, description, rating, image, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE title=?, description=?, rating=?, image=?, updated_at=?`,
		cake.ID, cake.Title, cake.Description, cake.Rating, cake.Image, createdAt, updatedAt,
		cake.Title, cake.Description, cake.Rating, cake.Image, updatedAt)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	cake.ID = uint(id)
	cake.CreatedAt = createdAt
	cake.UpdatedAt = updatedAt
	return cake, nil
}

func (r *cakeRepositoryImpl) DeleteCakeByID(ctx context.Context, cakeID uint) error {
	result, err := r.db.Exec("DELETE FROM cakes WHERE id = ?", cakeID)
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if affected == 0 {
		return core.ErrRecordNotFound
	}
	return nil
}
