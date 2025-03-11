package main

import (
	"fmt"
	"github.com/MoXcz/personal-webiste/internal/app"
)

func main() {
	srv := app.NewServer()

	fmt.Println("Server started on http://localhost:8080")
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
