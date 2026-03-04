package user

import (
	"sync"

	"github.com/google/uuid"
)

type Repository struct {
	mu     sync.RWMutex
	data   map[string]User
}

func NewRepository() *Repository {
	r := &Repository{
		data:   make(map[string]User),
	}

	r.Create("Igor")
	r.Create("Anna")

	return r
}

func (r *Repository) List() []User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]User, 0, len(r.data))
	for _, u := range r.data {
		out = append(out, u)
	}
	return out
}

func (r *Repository) GetByID(id string) (User, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.data[id]
	return u, ok
}

func (r *Repository) Create(name string) User {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uuid.New().String()
	u := User{ID: id, Name: name}
	r.data[id] = u

	return u
}

func (r *Repository) Delete(id string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[id]; !ok {
		return false
	}
	delete(r.data, id)
	return true
}