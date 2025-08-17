package db

import (
	m "mealmate/internal/model"
	w "mealmate/internal/warnings"
	"mealmate/pkg"
	"strings"
	"sync"
)

// Interface
type DBFooder interface {
	UpdateFood(string, map[string]any) (m.Food, w.Warning)
	DeleteFood(string) (m.Food, w.Warning)
	TakeFood(string) (m.Food, w.Warning)
	TakeFoodStore() map[string]*m.Food
	PutFood(*m.Food) w.Warning
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

func (d *DB) PutFood(f *m.Food) w.Warning {
	// Формируем ключ Food
	key := strings.TrimSpace(f.Name)

	// Check for inic
	isUnic := d.checkForUnic(key)
	if !isUnic {
		return notUnicFoodWarn
	}

	// Делаем новую запись с Food
	d.mu.Lock()
	d.FoodStore[key] = f
	d.mu.Unlock()
	return nil
}

func (d *DB) TakeFood(id string) (m.Food, w.Warning) {
	d.muR.RLock()
	defer d.muR.RUnlock()

	takingFood, ok := d.FoodStore[id]
	if !ok {
		return m.Food{}, notFoundFoodWarn
	}
	return *takingFood, nil
}

func (d *DB) UpdateFood(id string, forUpData map[string]any) (m.Food, w.Warning) {
	d.mu.Lock()
	defer d.mu.Unlock()

	takingFood, ok := d.FoodStore[id]
	if !ok {
		return m.Food{}, notFoundFoodWarn
	}
	takingFood = d.exReflecter.CrossUpdateStructs(takingFood, forUpData).(*m.Food)
	return *takingFood, nil
}

func (d *DB) DeleteFood(id string) (m.Food, w.Warning) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if deletingFood, ok := d.FoodStore[id]; ok {
		delete(d.FoodStore, id)
		return *deletingFood, nil
	}
	return m.Food{}, notFoundFoodWarn
}
