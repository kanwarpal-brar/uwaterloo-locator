package sqlite

import (
	"context"
	"database/sql"
	"math"

	"washroom-data-service/models"
	"washroom-data-service/repository"
)

type sqliteLocationQuery struct {
	db *sql.DB
}

func NewSQLiteLocationQuery(db *sql.DB) repository.LocationQueryRepository {
	return &sqliteLocationQuery{db: db}
}

func (r *sqliteLocationQuery) FindNearby(ctx context.Context, lat, lng, radius float64) ([]models.Washroom, error) {
	// SQLite doesn't have built-in geospatial functions, so we'll filter in memory
	query := `SELECT id, name, latitude, longitude, building, floor, 
	                 gender, is_accessible, version, created_at, updated_at
	          FROM washrooms`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var washrooms []models.Washroom
	for rows.Next() {
		w := models.Washroom{}
		err := rows.Scan(&w.ID, &w.Name, &w.Location.Latitude, &w.Location.Longitude,
			&w.Building, &w.Floor, &w.Gender, &w.IsAccessible,
			&w.Version, &w.CreatedAt, &w.UpdatedAt)
		if err != nil {
			return nil, err
		}

		dist := distance(lat, lng, w.Location.Latitude, w.Location.Longitude)
		if dist <= radius {
			washrooms = append(washrooms, w)
		}
	}
	return washrooms, rows.Err()
}

func (r *sqliteLocationQuery) FindInBuilding(ctx context.Context, building string) ([]models.Washroom, error) {
	query := `SELECT id, name, latitude, longitude, building, floor, 
	                 gender, is_accessible, version, created_at, updated_at
	          FROM washrooms WHERE building = ?`

	return r.queryWashrooms(ctx, query, building)
}

func (r *sqliteLocationQuery) FindByFloor(ctx context.Context, building string, floor int) ([]models.Washroom, error) {
	query := `SELECT id, name, latitude, longitude, building, floor, 
	                 gender, is_accessible, version, created_at, updated_at
	          FROM washrooms WHERE building = ? AND floor = ?`

	return r.queryWashrooms(ctx, query, building, floor)
}

func (r *sqliteLocationQuery) queryWashrooms(ctx context.Context, query string, args ...interface{}) ([]models.Washroom, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var washrooms []models.Washroom
	for rows.Next() {
		w := models.Washroom{}
		err := rows.Scan(&w.ID, &w.Name, &w.Location.Latitude, &w.Location.Longitude,
			&w.Building, &w.Floor, &w.Gender, &w.IsAccessible,
			&w.Version, &w.CreatedAt, &w.UpdatedAt)
		if err != nil {
			return nil, err
		}
		washrooms = append(washrooms, w)
	}
	return washrooms, rows.Err()
}

// distance calculates the distance between two points using the Haversine formula
func distance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Earth radius in meters
	φ1 := lat1 * math.Pi / 180
	φ2 := lat2 * math.Pi / 180
	Δφ := (lat2 - lat1) * math.Pi / 180
	Δλ := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
