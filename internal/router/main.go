package router

import (
	db "mealmate/internal/db"
	h "mealmate/internal/handler"
	s "mealmate/internal/service"

	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	// Handler for Food
	dB := db.NewDB()
	foodS := s.NewFoodServ(dB)
	foodH := h.NewFoodHandler(foodS)

	r.Route("/", func(r chi.Router) {

		r.Route("/food/", func(r chi.Router) {
			// r.Delete("/", food.Delete)
			// r.Patch("/", food.Update)
			r.Post("/", foodH.Create)
			// r.Get("/", food.Read)

			// TODO! Заглушка Для любых http - методов, которых нет
			// r.Handle("/", food)
		})
	})

	return r
}
