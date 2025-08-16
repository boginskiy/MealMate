package db

import (
	m "mealmate/internal/model"
	"mealmate/pkg"
	"strings"
	"sync"
)

// Interface
type DBFooder interface {
	ChangeFood(string, map[string]any) (m.Food, warning)
	TakeFood(string) (m.Food, warning)
	TakeFoodStore() map[string]*m.Food
	PutFood(*m.Food) warning
}

type DB struct {
	FoodStore   map[string]*m.Food
	exReflecter pkg.ExReflecter
	muR         sync.RWMutex
	mu          sync.Mutex
}

func NewDB(exref pkg.ExReflecter) *DB {
	return &DB{
		FoodStore:   make(map[string]*m.Food, 5),
		exReflecter: exref,
	}
}

func (d *DB) checkForUnic(id string) bool {
	d.muR.RLock()
	defer d.muR.RUnlock()

	if _, ok := d.FoodStore[id]; ok {
		return false
	}
	return true
}

func (d *DB) TakeFoodStore() map[string]*m.Food {
	return d.FoodStore
}

func (d *DB) PutFood(f *m.Food) warning {
	// Формируем ключ Food
	key := strings.TrimSpace(f.Name)

	// Check for inic
	isUnic := d.checkForUnic(key)
	if !isUnic {
		return unicFoodWarn
	}

	// Делаем новую запись с Food
	d.mu.Lock()
	d.FoodStore[key] = f
	d.mu.Unlock()
	return ""
}

func (d *DB) TakeFood(id string) (m.Food, warning) {
	d.muR.RLock()
	defer d.muR.RUnlock()

	takingFood, ok := d.FoodStore[id]
	if !ok {
		return m.Food{}, notFoundFoodWarn
	}
	return *takingFood, ""
}

func (d *DB) ChangeFood(id string, forUpData map[string]any) (m.Food, warning) {
	d.mu.Lock()
	defer d.mu.Unlock()

	takingFood, ok := d.FoodStore[id]
	if !ok {
		return m.Food{}, notFoundFoodWarn
	}
	takingFood = d.exReflecter.CrossUpdateStructs(takingFood, forUpData).(*m.Food)
	return *takingFood, ""
}
