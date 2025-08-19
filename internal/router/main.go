package router

import (
	a "mealmate/internal/alerts"
	db "mealmate/internal/db"
	h "mealmate/internal/handler"
	s "mealmate/internal/service"
	"mealmate/pkg"

	"github.com/go-chi/chi"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	// pkg
	exReflect := pkg.NewExtraReflect()
	exEncode := pkg.NewExtraEncode()

	// alerts
	alert := a.NewAlert()

	// Food Service
	foodStore := db.NewFoodStore()
	foodRepos := db.NewRep(exReflect, foodStore)

	foodS := s.NewFoodServ(exReflect, exEncode, alert, foodRepos)
	foodH := h.NewMealHandler(foodS)

	// Ingredient Service
	// ingredStore := db.NewIngredStore()
	// ingredRepos := db.NewRep(exReflect, ingredStore)

	ingredS := s.NewIngredServ()
	ingredH := h.NewMealHandler(ingredS)

	r.Route("/", func(r chi.Router) {

		r.Route("/food/", func(r chi.Router) {
			r.MethodNotAllowed(foodH.ServeHTTP)
			r.Delete("/", foodH.Delete)
			r.Patch("/", foodH.Update)
			r.Post("/", foodH.Create)
			r.Get("/", foodH.Read)
		})

		r.Route("/ingred/", func(r chi.Router) {
			r.MethodNotAllowed(foodH.ServeHTTP)
			r.Delete("/", ingredH.Delete)
			r.Patch("/", ingredH.Update)
			r.Post("/", ingredH.Create)
			r.Get("/", ingredH.Read)
		})
	})

	return r
}
