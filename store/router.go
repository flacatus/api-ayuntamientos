package store

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Method to get all townhalls",
		"GET",
		"/api/v1/townhalls/all",
		controller.GetAllTownHalls,
	}, {
		"Method to insert townhalls",
		"POST",
		"/api/v1/townhalls/insert",
		controller.InsertTHalls,
	},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
