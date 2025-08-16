package router

import (
	db "mealmate/internal/db"
	h "mealmate/internal/handler"
	s "mealmate/internal/service"
	"mealmate/pkg"

	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	// Handler for Food
	exReflect := pkg.NewExtraReflect()
	db := db.NewDB(exReflect)

	foodS := s.NewFoodServ(exReflect, db)
	foodH := h.NewFoodHandler(foodS)

	r.Route("/", func(r chi.Router) {

		r.Route("/food/", func(r chi.Router) {
			// r.Delete("/", food.Delete)
			r.Patch("/", foodH.Update)
			r.Post("/", foodH.Create)
			r.Get("/", foodH.Read)

			// TODO! Заглушка Для любых http - методов, которых нет
			// r.Handle("/", food)
		})
	})

	return r
}
