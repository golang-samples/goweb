package article

import (
	"github.com/stretchrcom/goweb/goweb"
	"labix.org/v2/mgo"
	"log"
)

const COLLECTION = "article"

func Init(db *mgo.Database) {
	log.Println("Initializing article package")
	controller := NewController(db)
	goweb.MapRest("/articles", controller)
}
