package api

import (
	"./article"
	"labix.org/v2/mgo"
	"log"
)

func Init(db *mgo.Database) {
	log.Println("Initializing api package")
	// article package
	article.Init(db)
}
