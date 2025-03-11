package app

import (
	"fmt"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", Home)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to My Go App</h1><p>This is a simple homepage.</p>")
}
