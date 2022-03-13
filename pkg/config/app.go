package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB //proměnná, která pomáhá souborům v projektu komunikovat s databází
)

func Connect() {
	d, err := gorm.Open("mysql", "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local")    //pokud spojení neprobjehne vypíše se err pokud probjehne spojení databáze se dostane do d (gorm.open() otevírá spojení s db, je třeba mu dát jméno db, jmeno uživatele a heslo
	if err != nil{ // pokud err není nil tzn. došlo k chybě pak panic()
		panic(err)
	}
	db = d
}//propojovací funkce, která nám pomáhá otevří databázi a spojit se s ní

func GetDB() *gorm.DB{ // funkce která vrátí databázi, takže když ji teď budu volat v jiných souborech mohu s databází komunikovat 

	return db
}