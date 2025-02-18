package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"
	"washroom-data-service/models"
)

// Initialize creates and initializes the SQLite database if it doesn't exist
func Initialize(dbPath string, loadTestData bool) (*sql.DB, error) {
	needsInit := !fileExists(dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if needsInit {
		if err := createSchema(db); err != nil {
			return nil, fmt.Errorf("error creating schema: %v", err)
		}

		if loadTestData {
			if err := loadBuildingsData(db); err != nil {
				return nil, fmt.Errorf("error loading test data: %v", err)
			}
		}
	}

	return db, nil
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func createSchema(db *sql.DB) error {
	schemaSQL, err := os.ReadFile("washroom-data-service/repository/sqlite/schema.sql")
	if err != nil {
		return err
	}

	if _, err := db.Exec(string(schemaSQL)); err != nil {
		return err
	}

	return nil
}

func loadBuildingsData(db *sql.DB) error {
	data, err := os.ReadFile("washroom-data-service/testdata/buildings.json")
	if err != nil {
		return err
	}

	var buildingsData struct {
		Features []struct {
			Properties struct {
				BuildingCode string  `json:"buildingCode"`
				BuildingName string  `json:"buildingName"`
				Latitude     float64 `json:"latitude"`
				Longitude    float64 `json:"longitude"`
			} `json:"properties"`
		} `json:"features"`
	}

	if err := json.Unmarshal(data, &buildingsData); err != nil {
		return err
	}

	// Convert each building into a washroom entry
	for _, feature := range buildingsData.Features {
		if feature.Properties.Latitude == 0 || feature.Properties.Longitude == 0 {
			continue // Skip buildings with no coordinates
		}

		washroom := models.NewWashroom()
		washroom.Name = fmt.Sprintf("%s Main Washroom", feature.Properties.BuildingCode)
		washroom.Location = models.Location{
			Latitude:  feature.Properties.Latitude,
			Longitude: feature.Properties.Longitude,
		}
		washroom.Building = feature.Properties.BuildingCode
		washroom.Floor = 1
		washroom.Gender = "All-Gender"
		washroom.IsAccessible = true
		washroom.CreatedAt = time.Now()
		washroom.UpdatedAt = time.Now()

		query := `
			INSERT INTO washrooms (
				name, latitude, longitude, building, floor, 
				gender, is_accessible, version, created_at, updated_at
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`
		_, err := db.Exec(query,
			washroom.Name, washroom.Location.Latitude, washroom.Location.Longitude,
			washroom.Building, washroom.Floor, washroom.Gender, washroom.IsAccessible,
			washroom.Version, washroom.CreatedAt, washroom.UpdatedAt)

		if err != nil {
			return fmt.Errorf("error inserting washroom: %v", err)
		}
	}

	return nil
}
