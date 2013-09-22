package article

import (
	"github.com/stretchr/goweb"
	"labix.org/v2/mgo"
	"log"
)

const COLLECTION = "article"

func Init(db *mgo.Database) {
	log.Println("Initializing article package")
	controller := NewController(db)
	goweb.MapController("/articles", controller)
}
