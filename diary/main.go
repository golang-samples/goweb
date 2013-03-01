package main

import (
	"./api"
	"flag"
	"github.com/stretchrcom/goweb/goweb"
	"labix.org/v2/mgo"
	"log"
	"net/http"
)

var (
	addr   string
	dbhost string
	dbname string
)

func init() {
	flag.StringVar(&addr, "http", ":8080", "Addres of web server")
	flag.StringVar(&dbhost, "dbhost", "localhost", "DB hosts")
	flag.StringVar(&dbname, "dbname", "diary", "DB name")
}

func main() {

	flag.Parse()

	log.Println("Accessing DB...")
	session, err := mgo.Dial(dbhost)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(dbname)
	log.Printf("DB (%s %s) opend.", dbhost, dbname)

	log.Println("Initializing handlers...")
	api.Init(db)
	goweb.ConfigureDefaultFormatters()
	http.Handle("/api/", http.StripPrefix("/api", goweb.DefaultHttpHandler))
	log.Println("Registered a handler for RESTful API.")

	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Println("Registered a handler for static files.")

	log.Printf("Starting Diary Server at %s", addr)
	http.ListenAndServe(addr, nil)
}
