package routes

import (
	"github.com/Fi44er/go-projects/tree/main/3_go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStore = func(router *mux.Router) {
	router.Handle("/book/", controllers.CreateBook)
}
