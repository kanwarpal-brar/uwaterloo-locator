package memory

import (
	"context"
	"sync"
	"time"

	"github.com/brark/uwaterloo-locator/washroom-data-service/models"
	"github.com/brark/uwaterloo-locator/washroom-data-service/repository"
)

type memoryRepository struct {
	mutex     sync.RWMutex
	washrooms map[string]models.Washroom
}

func NewMemoryRepository() repository.WashroomRepository {
	return &memoryRepository{
		washrooms: make(map[string]models.Washroom),
	}
}

func (r *memoryRepository) Create(ctx context.Context, w *models.Washroom) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	w.CreatedAt = time.Now()
	w.UpdatedAt = time.Now()
	r.washrooms[w.ID] = *w
	return nil
}

func (r *memoryRepository) GetByID(ctx context.Context, id string) (*models.Washroom, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if w, exists := r.washrooms[id]; exists {
		return &w, nil
	}
	return nil, repository.ErrNotFound
}

func (r *memoryRepository) Update(ctx context.Context, w *models.Washroom) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.washrooms[w.ID]; !exists {
		return repository.ErrNotFound
	}

	w.UpdatedAt = time.Now()
	r.washrooms[w.ID] = *w
	return nil
}

func (r *memoryRepository) Delete(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.washrooms[id]; !exists {
		return repository.ErrNotFound
	}

	delete(r.washrooms, id)
	return nil
}
