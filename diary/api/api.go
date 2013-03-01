package api

import (
	"labix.org/v2/mgo"
	"log"
	"./article"
)

func Init(db *mgo.Database) {
	log.Println("Initializing api package")
	// article package
	article.Init(db)
}
