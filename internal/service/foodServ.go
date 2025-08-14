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
		tmpErr := m.NewErrorWarn(http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req, "error"), http.StatusBadRequest
	}

	// Deserialization
	newFood := m.NewFood()
	err = json.Unmarshal(tmpByte, newFood)

	if err != nil {
		tmpErr := m.NewErrorWarn(http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req, "error"), http.StatusBadRequest
	}

	// Write in DB
	war := f.db.PutFood(newFood)
	if war != "" {
		tmpWar := m.NewErrorWarn(http.StatusBadRequest, string(war), req.URL.Path)
		return tmpWar.PreparBody(req, "warning"), http.StatusUnprocessableEntity
	}

	return tmpByte, http.StatusOK
}
