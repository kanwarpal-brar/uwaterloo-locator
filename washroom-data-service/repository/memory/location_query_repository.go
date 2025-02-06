package memory

import (
	"context"
	"math"

	"github.com/brark/uwaterloo-locator/washroom-data-service/models"
	"github.com/brark/uwaterloo-locator/washroom-data-service/repository"
)

type memoryLocationQuery struct {
	washroomRepo *memoryRepository
}

func NewMemoryLocationQuery(repo *memoryRepository) repository.LocationQueryRepository {
	return &memoryLocationQuery{
		washroomRepo: repo,
	}
}

func (r *memoryLocationQuery) FindNearby(ctx context.Context, lat, lng, radius float64) ([]models.Washroom, error) {
	r.washroomRepo.mutex.RLock()
	defer r.washroomRepo.mutex.RUnlock()

	var result []models.Washroom
	for _, w := range r.washroomRepo.washrooms {
		if distance(lat, lng, w.Location.Latitude, w.Location.Longitude) <= radius {
			result = append(result, w)
		}
	}
	return result, nil
}

func (r *memoryLocationQuery) FindInBuilding(ctx context.Context, building string) ([]models.Washroom, error) {
	r.washroomRepo.mutex.RLock()
	defer r.washroomRepo.mutex.RUnlock()

	var result []models.Washroom
	for _, w := range r.washroomRepo.washrooms {
		if w.Building == building {
			result = append(result, w)
		}
	}
	return result, nil
}

func (r *memoryLocationQuery) FindByFloor(ctx context.Context, building string, floor int) ([]models.Washroom, error) {
	r.washroomRepo.mutex.RLock()
	defer r.washroomRepo.mutex.RUnlock()

	var result []models.Washroom
	for _, w := range r.washroomRepo.washrooms {
		if w.Building == building && w.Floor == floor {
			result = append(result, w)
		}
	}
	return result, nil
}

// distance returns the distance in meters between two points using the Haversine formula
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
