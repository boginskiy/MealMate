package service

import (
	"fmt"
	"mealmate/internal/db"
	m "mealmate/internal/model"
	w "mealmate/internal/warnings"
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

func (f *FoodServ) handleAlert(req *http.Request, msg any, statusCode int) ([]byte, int) {
	var message string
	var kind string

	switch v := msg.(type) {
	case error:
		message = v.Error()
		kind = "error"
	case w.Warning:
		message = v.Warning()
		kind = "warning"
	default:
		message = fmt.Sprintf("%v\n", msg)
		kind = "unknown"
	}

	alert := m.NewAlert(kind, statusCode, message, req.URL.Path)
	body := alert.PreparBody(req)
	return body, statusCode
}

func (f *FoodServ) getFoodID(deserializedFood map[string]any) (string, error) {
	tmpId, ok := deserializedFood[ID]
	if !ok {
		return "", notIdFieldErr
	}
	id, ok := tmpId.(string)
	if !ok {
		return "", notValidIdErr
	}
	return id, nil
}

func (f *FoodServ) Create(req *http.Request) ([]byte, int) {
	// Read Body
	body, err := f.exFuncer.ReadRequestBody(req)
	if err != nil {
		return f.handleAlert(req, err, http.StatusBadRequest)
	}

	// Deserialization
	newFood := m.NewFood()
	err = f.exEncoder.Deserialization(body, newFood)
	if err != nil {
		return f.handleAlert(req, err, http.StatusBadRequest)
	}

	// Write in DB
	warn := f.dBer.PutFood(newFood)
	if warn != nil {
		return f.handleAlert(req, warn, http.StatusUnprocessableEntity)
	}
	return body, http.StatusOK
}

/*
Params:

	?id=2              # Идентификатор блюда (по умолчанию: просмотр всех блюд)
	?category=Fastfood # Категория блюда (пример: Fastfood)
	?name=Burger       # Название блюда (пример: Burger)
	?totalPrice=10.25  # Стоимость блюда (пример: 10.25)
	?type=Vegetarian   # Тип блюда (пример: Vegetarian)

Default behavior: View all available foods.
Todo: Set up pagination.
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
	tmpByte, err := f.exEncoder.Serialization(tmpStore)
	if err != nil {
		return f.handleAlert(req, err, http.StatusBadRequest)
	}
	return tmpByte, http.StatusOK
}

func (f *FoodServ) Update(req *http.Request) ([]byte, int) {
	// Read Body
	body, err := f.exFuncer.ReadRequestBody(req)
	if err != nil {
		return f.handleAlert(req, err, http.StatusBadRequest)
	}

	// Deserialization to map
	var foodForUpdate map[string]any
	err = f.exEncoder.Deserialization(body, &foodForUpdate)
	if err != nil {
		return f.handleAlert(req, err, http.StatusBadRequest)
	}

	// Check that is ID, like 'Name'
	id, err := f.getFoodID(foodForUpdate)
	if err != nil {
		return f.handleAlert(req, err, http.StatusBadRequest)
	}

	// Give updating newFood into db
	updatedFood, warn := f.dBer.UpdateFood(id, foodForUpdate)
	if warn != nil {
		return f.handleAlert(req, warn, http.StatusBadRequest)
	}

	// Serialization
	tmpByte, err := f.exEncoder.Serialization(updatedFood)
	if err != nil {
		return f.handleAlert(req, err, http.StatusBadRequest)
	}
	return tmpByte, http.StatusOK
}

func (f *FoodServ) Delete(req *http.Request) ([]byte, int) {
	// Пробуем достать ID == Name
	nameFood := f.exFuncer.TakeIDFromPath(req, ID)

	// Пробуем достать ID из Body
	if nameFood == "" {
		body, err := f.exFuncer.ReadRequestBody(req)
		if err != nil {
			return f.handleAlert(req, err, http.StatusBadRequest)
		}

		var foodForDelete map[string]any
		err = f.exEncoder.Deserialization(body, &foodForDelete)
		if err != nil {
			return f.handleAlert(req, err, http.StatusBadRequest)
		}

		// Check that is ID, like 'Name'
		nameFood, err = f.getFoodID(foodForDelete)
		if err != nil {
			return f.handleAlert(req, err, http.StatusBadRequest)
		}
	}

	// Delete Food into db
	deletedFood, warn := f.dBer.DeleteFood(nameFood)
	if warn != nil {
		return f.handleAlert(req, warn, http.StatusBadRequest)
	}

	// Serialization
	tmpByte, err := f.exEncoder.Serialization(deletedFood)
	if err != nil {
		f.handleAlert(req, err, http.StatusBadRequest)
	}
	return tmpByte, http.StatusOK
}

// Рефачим, проверяем, add Args
// Тесты надо
