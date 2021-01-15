package router

import "github.com/gorilla/mux"

// Generates a new router with configured routes
func Generate() *mux.Router {
	return mux.NewRouter()
}
