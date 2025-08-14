package service

import (
	"encoding/json"
	"fmt"
	"io"
	"mealmate/internal/db"
	m "mealmate/internal/model"
	"mealmate/pkg"
	"net/http"
	"os"
	"reflect"
	"strconv"
)

type FoodServer interface {
	CreateFood(*http.Request) ([]byte, int)
	ReadFood(*http.Request) ([]byte, int)
}

type FoodServ struct {
	toolsStructer pkg.ToolsStructer
	dber          db.DBFooder
}

func NewFoodServ(toolsS pkg.ToolsStructer, db db.DBFooder) *FoodServ {
	return &FoodServ{toolsStructer: toolsS, dber: db}
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
	war := f.dber.PutFood(newFood)
	if war != "" {
		tmpWar := m.NewErrorWarn("warning", http.StatusUnprocessableEntity, string(war), req.URL.Path)
		return tmpWar.PreparBody(req), http.StatusUnprocessableEntity
	}

	return tmpByte, http.StatusOK
}

/*
Params:

	default - view all food
	extra - настроить пагинацию
*/
func (f *FoodServ) ReadFood(req *http.Request) ([]byte, int) {
	// tmpStore
	tmpStore := make([]m.Food, 0, 10)

	// Params of URL
	queryParams := req.URL.Query()

	for i, food := range f.dber.TakeFoodStore() {
		// Filter about queryParams
	}

	// TODO
	// Дальше отбор по параметрам сущностей. Вывод
	// Логику разнести по своим местам

	// Get fields of struct Food
	_type := reflect.TypeOf(*m.NewFood())
	numFields := _type.NumField()

	// Проходим по каждому полю структуры
	for i := 0; i < numFields; i++ {
		field := _type.Field(i)
		name := field.Name
		kind := field.Type.Kind()

		paramStr := queryParams.Get(name)

		// Разбор типов полей
		switch kind {
		case reflect.Int:
			paramInt, err := strconv.Atoi(paramStr)

			if err != nil {
				fmt.Fprintf(os.Stdout, "%v\n", err)
				continue
			}

		case reflect.String:
			paramStr
		case reflect.Float64:
			paramFlt
		}
	}

	// fieldsList := f.toolsStructer.ShowdownFields(*)
	// for _, field := range fieldsList {
	// 	queryParams.Get(field)
	// }

	return []byte{}, 500

}
