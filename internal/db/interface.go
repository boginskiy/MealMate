package db

import (
	w "mealmate/internal/alerts/warnings"
	m "mealmate/internal/model"
)

type FStorer interface {
	getRecord(string) (*m.FoodModel, bool)
	setRecord(string, m.Modeler) bool
	getStore() map[string]*m.FoodModel
	getEmptyRecord() m.FoodModel
	deleteRecord(string)
}

type IStorer interface {
	getRecord(string) (*m.IngredModel, bool)
	setRecord(string, *m.IngredModel) bool
	getStore() map[string]*m.IngredModel
	getEmptyRecord() m.IngredModel
	deleteRecord(string)
}

type Repository interface {
	UpdateRecord(string, map[string]any) (m.Modeler, w.Warning)
	DeleteRecord(string) (m.Modeler, w.Warning)
	TakeRecord(string) (m.Modeler, w.Warning)
	TakeAllStore() map[string]*m.FoodModel // Вынести отдельн ?
	PutRecord(m.Modeler) w.Warning
}
