package routes

import (
	"github.com/gorilla/mux"
	"github.com/libkarl/go-bookstore/pkg/controllers" //go mod edit -module github.com/user/a-different-repo
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
}
