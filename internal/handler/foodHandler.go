package handler

import (
	"fmt"

	s "mealmate/internal/service"
	"net/http"
	"os"
)

type FoodHandler struct {
	foodServer s.FoodServer
}

func NewFoodHandler(f s.FoodServer) *FoodHandler {
	return &FoodHandler{foodServer: f}
}

func (f *FoodHandler) Create(res http.ResponseWriter, req *http.Request) {
	body, status := f.foodServer.CreateFood(req)
	// Response
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(body)
}

func (f *FoodHandler) Read(res http.ResponseWriter, req *http.Request) {
	body, status := f.foodServer.ReadFood(req)
	// Response
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(body)
}

// ServeHTTP Базовый обработчик для FoodHandler
func (f *FoodHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusNotFound)

	_, err := fmt.Fprint(res, "Oops! We couldn't find anything here")
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
	}
}
