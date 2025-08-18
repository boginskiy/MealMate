package db

import m "mealmate/internal/model"

type FoodStore struct {
	Store map[string]*m.FoodModel // view table for Food
}

func NewFoodStore() *FoodStore {
	return &FoodStore{make(map[string]*m.FoodModel, 5)}
}

func (f *FoodStore) getEmptyRecord() m.FoodModel {
	return m.FoodModel{}
}

func (f *FoodStore) getRecord(id string) (*m.FoodModel, bool) {
	value, ok := f.Store[id]
	return value, ok
}

func (f *FoodStore) getStore() map[string]*m.FoodModel {
	return f.Store
}

func (f *FoodStore) setRecord(id string, value m.Modeler) (ok bool) {
	f.Store[id] = value.(*m.FoodModel)
	return true
}

func (f *FoodStore) deleteRecord(id string) {
	delete(f.Store, id)
}
