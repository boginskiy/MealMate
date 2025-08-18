package db

import (
	w "mealmate/internal/alerts/warnings"
	m "mealmate/internal/model"
	"mealmate/pkg"
	"strings"
	"sync"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type FoodRepo struct {
	exRefl  pkg.ExReflecter
	muR     sync.RWMutex
	mu      sync.Mutex
	FStorer FStorer
}

func NewFoodRepo(exrefl pkg.ExReflecter, fstore FStorer) *FoodRepo {
	return &FoodRepo{exRefl: exrefl, FStorer: fstore}
}

func (f *FoodRepo) checkForUnic(id string) bool {
	f.muR.RLock()
	defer f.muR.RUnlock()
	if _, ok := f.FStorer.getRecord(id); ok {
		return false
	}
	return true
}

func (f *FoodRepo) TakeAllStore() map[string]*m.FoodModel {
	return f.FStorer.getStore()
}

func (f *FoodRepo) PutRecord(value m.Modeler) w.Warning {
	// Формируем ключ Food
	key := strings.TrimSpace(value.GetName())
	// Check for inic
	isUnic := f.checkForUnic(key)
	if !isUnic {
		return notUnicFoodWarn
	}
	// Делаем новую запись с Food
	f.mu.Lock()
	f.FStorer.setRecord(key, value.(*m.FoodModel))
	f.mu.Unlock()
	return nil
}

func (f *FoodRepo) TakeRecord(id string) (m.Modeler, w.Warning) {
	f.muR.RLock()
	defer f.muR.RUnlock()
	tookRecord, ok := f.FStorer.getRecord(id)
	if !ok {
		return f.FStorer.getEmptyRecord(), notFoundFoodWarn
	}
	return *tookRecord, nil
}

func (f *FoodRepo) UpdateRecord(id string, partForUpdate map[string]any) (m.Modeler, w.Warning) {
	f.mu.Lock()
	defer f.mu.Unlock()

	tookRecord, ok := f.FStorer.getRecord(id)
	if !ok {
		return f.FStorer.getEmptyRecord(), notFoundFoodWarn
	}
	// Update old Record on new
	tookRecord = f.exRefl.CrossUpdateStructs(tookRecord, partForUpdate).(*m.FoodModel)
	return *tookRecord, nil
}

func (f *FoodRepo) DeleteRecord(id string) (m.Modeler, w.Warning) {
	// TODO. Временно так
	titleCase := cases.Title(language.Russian)
	id = titleCase.String(id)

	f.mu.Lock()
	defer f.mu.Unlock()

	if deletingRecord, ok := f.FStorer.getRecord(id); ok {
		f.FStorer.deleteRecord(id)
		return *deletingRecord, nil
	}
	return f.FStorer.getEmptyRecord(), notFoundFoodWarn
}
