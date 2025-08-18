package server

import (
	"fmt"
	c "mealmate/cmd/config"
	r "mealmate/internal/router"
	"net/http"
	"os"
)

func Run() error {
	// agrs - атрибуты командной строки
	argsCLI := c.NewArgsCLI()

	fmt.Fprintf(os.Stdout, "'MealMate' server has started successfully on port %v\n", argsCLI.Port)
	return http.ListenAndServe(argsCLI.Port, r.Router())
}
