 package models

 import (
	"github.com/jinzhu/gorm"
	"github.com/libkarl/go-bookstore/pkg/config"
 )

 var db *gorm.DB

 type Book struct { //struktura knihy, k čemu je tam ten gorm.model?
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`	
 }

func init(){ // funkce pro inicializaci databáze
	config.Connect() // funkce kterou jsme vytvořili v App.go, vytvoří spojení a pokud se podaří výsledek se předá do GetDB
	db = config.GetDB()
	db.AutoMigrate(&Book{}) // vytváří tabulky, indexy  apod. mění existující data v tabulce pokud se změní, ale neodstaňuje nepoužívané data, aby ochráníla data v uložaná v databázi
}

// program funguje tak, že kontrolery v cestách (routes), předávají kontrolu book.go, což
// je správný model, tok dat funguje v tom smyslu, že uživatel interaguje s trasami
// a trasi posílají ovládání kontrolerům, kde bude veškerá logika která bude směřovat do book.go
// které řídí operace  s databází
