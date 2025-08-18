package db

import m "mealmate/internal/model"

type IngredStore struct {
	Store map[string]*m.IngredModel // view table Ingredient
}

func NewIngredStore() *IngredStore {
	return &IngredStore{make(map[string]*m.IngredModel, 5)}
}

func (i *IngredStore) getEmptyRecord() m.IngredModel {
	return m.IngredModel{}
}

func (i *IngredStore) getRecord(id string) (*m.IngredModel, bool) {
	value, ok := i.Store[id]
	return value, ok
}

func (i *IngredStore) getStore() map[string]*m.IngredModel {
	return i.Store
}

func (i *IngredStore) setRecord(id string, value m.Modeler) (ok bool) {
	i.Store[id] = value.(*m.IngredModel)
	return true
}

func (i *IngredStore) deleteRecord(id string) {
	delete(i.Store, id)
}
