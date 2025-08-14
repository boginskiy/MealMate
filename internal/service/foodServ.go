package service

import (
	"encoding/json"
	"io"
	db "mealmate/internal/db"
	m "mealmate/internal/model"
	"net/http"
)

type FoodServer interface {
	CreateFood(*http.Request) ([]byte, int)
}

type FoodServ struct {
	db db.DBFooder
}

func NewFoodServ(dB db.DBFooder) *FoodServ {
	return &FoodServ{db: dB}
}

func (f *FoodServ) CreateFood(req *http.Request) ([]byte, int) {
	// Read Body
	tmpByte, err := io.ReadAll(req.Body)

	// Check error
	if err != nil {
		tmpErr := m.NewErrorWarn("error", http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	// Deserialization
	newFood := m.NewFood()
	err = json.Unmarshal(tmpByte, newFood)

	if err != nil {
		tmpErr := m.NewErrorWarn("error", http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	// Write in DB
	war := f.db.PutFood(newFood)
	if war != "" {
		tmpWar := m.NewErrorWarn("warning", http.StatusUnprocessableEntity, string(war), req.URL.Path)
		return tmpWar.PreparBody(req), http.StatusUnprocessableEntity
	}

	return tmpByte, http.StatusOK
}
