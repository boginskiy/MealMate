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

func (f *FoodHandler) sendResponse(res http.ResponseWriter, body []byte, status int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	_, err := res.Write(body)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Problem in FoodHandler.sendResponse: %v\n", err)
	}
}

func (f *FoodHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Response
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusNotFound)
	_, err := fmt.Fprint(res, "Oops! We couldn't find anything here")
	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err)
	}
}

func (f *FoodHandler) Create(res http.ResponseWriter, req *http.Request) {
	body, status := f.servicer.Create(req)
	f.sendResponse(res, body, status)
}

func (f *FoodHandler) Read(res http.ResponseWriter, req *http.Request) {
	body, status := f.servicer.Read(req)
	f.sendResponse(res, body, status)
}

func (f *FoodHandler) Update(res http.ResponseWriter, req *http.Request) {
	body, status := f.servicer.Update(req)
	f.sendResponse(res, body, status)
}

func (f *FoodHandler) Delete(res http.ResponseWriter, req *http.Request) {
	body, status := f.servicer.Delete(req)
	f.sendResponse(res, body, status)
}
