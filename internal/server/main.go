package server

import (
	"fmt"
	r "mealmate/internal/router"
	"net/http"
	"os"
)

func Run() error {
	// TODO: Вывести в аргументы запуска приложения
	port := ":8080"

	fmt.Fprintf(os.Stdout, "'MealMate' server has started successfully on port %v\n", port)
	return http.ListenAndServe(port, r.Router())
}
