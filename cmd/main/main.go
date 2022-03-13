// zde jsou všechny funkce pro různé operace s databází propojené s kontrolery 
// také je tu vytvořená serverová strana programu a definice našeho local host 

package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/libkarl/go-bookstore/pkg/routes"
)

func main(){
	r := mux.NewRouter() //inicializace směrovače
	routes.RegisterBookStoreRoutes(r) //název proměnné v routes, kde jsou uložené trasy, předáme r  
}