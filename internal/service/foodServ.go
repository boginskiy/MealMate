package service

import (
	"encoding/json"
	"io"
	"mealmate/internal/db"
	m "mealmate/internal/model"
	"mealmate/pkg"
	"net/http"
)

const ID = "Name"

type FoodServ struct {
	exReflecter pkg.ExReflecter
	exEncoder   pkg.ExEncoder
	exFuncer    ExFuncer
	dBer        db.DBFooder
}

func NewFoodServ(exref pkg.ExReflecter, exencd pkg.ExEncoder, db db.DBFooder) *FoodServ {
	exfunc := NewExtraFunc()
	return &FoodServ{
		exReflecter: exref,
		exEncoder:   exencd,
		exFuncer:    exfunc,
		dBer:        db}
}

func (f *FoodServ) sendAlert(req *http.Request, kind, descrip string, status int) ([]byte, int) {
	tmpAl := m.NewAlert(kind, status, descrip, req.URL.Path)
	body := tmpAl.PreparBody(req)
	return body, status
}

func (f *FoodServ) Create(req *http.Request) ([]byte, int) {
	// Read Body
	tmpByte, err := io.ReadAll(req.Body)

	if err != nil {
		return f.sendAlert(req, "error", err.Error(), http.StatusBadRequest)
	}

	// Deserialization
	newFood := m.NewFood()
	err = f.exEncoder.Deserialization(tmpByte, tmpByte)

	if err != nil {
		return f.sendAlert(req, "error", err.Error(), http.StatusBadRequest)
	}

	// Write in DB
	war := f.dBer.PutFood(newFood)
	if war != "" {
		return f.sendAlert(req, "warning", string(war), http.StatusUnprocessableEntity)
	}

	return tmpByte, http.StatusOK
}

/*
Params:

	?id=2
	&
	?name=Burger
	&
	?type=Vegetarian
	&
	?category=Fastfood
	&
	?totalPrice=10.25

	default - view all food
	todo    - set up pagination
*/
func (f *FoodServ) Read(req *http.Request) ([]byte, int) {
	tmpStore := make([]m.Food, 0, 10) // tmpStore
	queryParams := req.URL.Query()    // Params of URL

	for _, food := range f.dBer.TakeFoodStore() {
		// Filter about queryParams
		if ok := f.exFuncer.NeedShow(*food, queryParams); ok {
			tmpStore = append(tmpStore, *food)
		}
	}

	// Serialization
	tmpByte, err := json.Marshal(tmpStore)
	if err != nil {
		tmpErr := m.NewAlert("error", http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	return tmpByte, http.StatusOK
}

func (f *FoodServ) Update(req *http.Request) ([]byte, int) {
	// Read Body
	tmpByte, err := io.ReadAll(req.Body)

	// Check error
	if err != nil {
		tmpErr := m.NewAlert("error", http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	// Deserialization to map
	var forUpFood map[string]any
	err = json.Unmarshal(tmpByte, &forUpFood)

	if err != nil {
		tmpErr := m.NewAlert("error", http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	// Check that is ID, like 'Name'
	tmpId, ok := forUpFood[ID]
	id, ok2 := tmpId.(string)

	if !ok || !ok2 {
		tmpErr := m.NewAlert("error", http.StatusBadRequest, badBodyErr.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	newFood, war := f.dBer.ChangeFood(id, forUpFood)
	if war != "" {
		tmpWar := m.NewAlert("warning", http.StatusBadRequest, string(war), req.URL.Path)
		return tmpWar.PreparBody(req), http.StatusBadRequest
	}

	// Serialization
	tmpByte2, err := json.Marshal(newFood)
	if err != nil {
		tmpErr := m.NewAlert("error", http.StatusBadRequest, err.Error(), req.URL.Path)
		return tmpErr.PreparBody(req), http.StatusBadRequest
	}

	return tmpByte2, http.StatusOK
}

func (f *FoodServ) Delete(req *http.Request) ([]byte, int) {
	return []byte{}, 0
}
