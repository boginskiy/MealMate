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

	exReflect := pkg.NewExtraReflect()
	exEncode := pkg.NewExtraEncode()
	db := db.NewDB(exReflect)

	foodS := s.NewFoodServ(exReflect, exEncode, db)
	foodH := h.NewFoodHandler(foodS)

	r.Route("/", func(r chi.Router) {

		r.Route("/food/", func(r chi.Router) {
			r.MethodNotAllowed(foodH.ServeHTTP)
			r.Delete("/", foodH.Delete)
			r.Patch("/", foodH.Update)
			r.Post("/", foodH.Create)
			r.Get("/", foodH.Read)
		})
	})

	return r
}
