package handler

import (
	"fmt"

	s "mealmate/internal/service"
	"net/http"
	"os"
)

type FoodHandler struct {
	servicer s.Servicer
}

func NewFoodHandler(s s.Servicer) *FoodHandler {
	return &FoodHandler{servicer: s}
}

func (f *FoodHandler) Create(res http.ResponseWriter, req *http.Request) {
	body, status := f.servicer.Create(req)
	// Response
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(body)
}

func (f *FoodHandler) Read(res http.ResponseWriter, req *http.Request) {
	body, status := f.servicer.Read(req)
	// Response
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(body)
}

func (f *FoodHandler) Update(res http.ResponseWriter, req *http.Request) {
	body, status := f.servicer.Update(req)
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
