package sqlite

import (
	"context"
	"database/sql"
	"time"

	"washroom-data-service/models"
	"washroom-data-service/repository"
)

type sqliteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) repository.WashroomRepository {
	return &sqliteRepository{db: db}
}

func (r *sqliteRepository) Create(ctx context.Context, w *models.Washroom) error {
	query := `
		INSERT INTO washrooms (
			name, latitude, longitude, building, floor, gender, 
			is_accessible, version, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	now := time.Now()
	w.CreatedAt = now
	w.UpdatedAt = now

	result, err := r.db.ExecContext(ctx, query,
		w.Name, w.Location.Latitude, w.Location.Longitude,
		w.Building, w.Floor, w.Gender, w.IsAccessible,
		w.Version, w.CreatedAt, w.UpdatedAt)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	w.ID = string(id)
	return nil
}

func (r *sqliteRepository) GetByID(ctx context.Context, id string) (*models.Washroom, error) {
	query := `
		SELECT id, name, latitude, longitude, building, floor, 
		       gender, is_accessible, version, created_at, updated_at
		FROM washrooms WHERE id = ?
	`
	w := &models.Washroom{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&w.ID, &w.Name, &w.Location.Latitude, &w.Location.Longitude,
		&w.Building, &w.Floor, &w.Gender, &w.IsAccessible,
		&w.Version, &w.CreatedAt, &w.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, repository.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (r *sqliteRepository) Update(ctx context.Context, w *models.Washroom) error {
	query := `
		UPDATE washrooms SET 
			name = ?, latitude = ?, longitude = ?, building = ?,
			floor = ?, gender = ?, is_accessible = ?, version = ?,
			updated_at = ?
		WHERE id = ?
	`
	w.UpdatedAt = time.Now()
	result, err := r.db.ExecContext(ctx, query,
		w.Name, w.Location.Latitude, w.Location.Longitude,
		w.Building, w.Floor, w.Gender, w.IsAccessible,
		w.Version, w.UpdatedAt, w.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}

func (r *sqliteRepository) Delete(ctx context.Context, id string) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM washrooms WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return repository.ErrNotFound
	}
	return nil
}
