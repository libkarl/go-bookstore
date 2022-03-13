package routes

import (
	"github.com/gorilla/mux"
	"github.com/libkarl/go-bookstore/pkg/controllers" //go mod edit -module github.com/user/a-different-repo
)

var RegisterBookStoreRoutes = func(router *mux.Router) { //trasy z backendu do databáze a zpět

	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST") //přidání knihy do databáze
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET") //vytáhnutí knihy z databáze
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET") //vytáhnutí knihy z databáze podle jejího ID
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT") // 
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE") // odstranění knihy z databáze
}