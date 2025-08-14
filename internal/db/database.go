package db

import (
	m "mealmate/internal/model"
	"strings"
	"sync"
)

// Interface
type DBFooder interface {
	PutFood(*m.Food) warning
	TakeFood(string) (*m.Food, warning)
	TakeFoodStore() map[string]*m.Food
}

type DB struct {
	FoodStore map[string]*m.Food
	muR       sync.RWMutex
	mu        sync.Mutex
}

func NewDB() *DB {
	return &DB{
		FoodStore: make(map[string]*m.Food, 5),
	}
}

func (d *DB) TakeFoodStore() map[string]*m.Food {
	return d.FoodStore
}

func (d *DB) PutFood(f *m.Food) warning {
	// Формируем ключ Food
	key := strings.TrimSpace(f.Name) + " " + strings.TrimSpace(f.Type)

	// Считываем Map по ключу. Проверка, что записывемый Food будет уникальным.
	d.muR.RLock()
	if _, ok := d.FoodStore[key]; ok {
		d.muR.RUnlock()
		return unicFoodWarn
	}
	d.muR.RUnlock()

	// Делаем новую запись с Food
	d.mu.Lock()
	d.FoodStore[key] = f
	d.mu.Unlock()
	return ""
}

func (d *DB) TakeFood(key string) (*m.Food, warning) {
	return nil, ""
}
