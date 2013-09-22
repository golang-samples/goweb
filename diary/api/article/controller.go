package article

import (
	"encoding/json"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

type Controller struct {
	db *mgo.Database
}

func jsonRequestDecoder(cx context.Context, v interface{}) error {
	data, err := cx.RequestBody()

	if err != nil {
		return err
	}

	return json.Unmarshal(data, &v)
}

func NewController(db *mgo.Database) *Controller {
	if db == nil {
		panic("database cannot be nil")
	}

	return &Controller{db}
}

func (cr *Controller) Create(cx context.Context) error {
	log.Println("Creating a article...")
	c := cr.db.C(COLLECTION)

	var article Article
	jsonRequestDecoder(cx, &article)

	article.Id = bson.NewObjectId().Hex()

	if err := c.Insert(&article); err != nil {
		log.Println("Error: %s", err.Error())
		return goweb.Respond.WithStatus(cx, http.StatusForbidden)
	}

	log.Printf("Created diary id=%s", article.Id)
	return goweb.API.RespondWithData(cx, article)
}

func (cr *Controller) Delete(id string, cx context.Context) error {
	log.Printf("Deleting a article id=%s...", id)
	c := cr.db.C(COLLECTION)
	if err := c.RemoveId(id); err != nil {
		log.Printf("Error: %s", err.Error())
		return goweb.Respond.WithStatus(cx, http.StatusForbidden)
	}

	log.Printf("Deleted article id=%s", id)
	return goweb.Respond.WithOK(cx)
}

func (cr *Controller) DeleteMany(cx context.Context) error {
	log.Println("Deleting all articles...")
	c := cr.db.C(COLLECTION)
	if _, err := c.RemoveAll(nil); err != nil {
		log.Println("Error: %s", err.Error())
		return goweb.Respond.WithStatus(cx, http.StatusForbidden)
	}

	log.Println("Deleted all articles")
	return goweb.Respond.WithOK(cx)
}

func (cr *Controller) Read(id string, cx context.Context) error {
	log.Printf("Read a article id=%s", id)
	c := cr.db.C(COLLECTION)
	var article Article
	if err := c.FindId(id).One(&article); err != nil {
		log.Println("Error: %s", err.Error())
		return goweb.Respond.WithStatus(cx, http.StatusForbidden)
	}

	log.Printf("Read article id=%s", id)
	return goweb.API.RespondWithData(cx, article)
}

func (cr *Controller) ReadMany(cx context.Context) error {
	log.Println("Read all articles...")
	c := cr.db.C(COLLECTION)
	count, err := c.Count()
	if err != nil {
		log.Println("Error: %s", err.Error())
		return goweb.Respond.WithStatus(cx, http.StatusForbidden)
	}

	articles := make([]*Article, count)
	if err := c.Find(nil).All(&articles); err != nil {
		log.Println("Error: %s", err.Error())
		return goweb.Respond.WithStatus(cx, http.StatusForbidden)
	}

	log.Printf("Read all %d articles", count)
	return goweb.API.RespondWithData(cx, articles)
}

func (cr *Controller) Replace(id string, cx context.Context) error {
	log.Printf("Update a article id=%s...", id)
	c := cr.db.C(COLLECTION)

	var article *Article
	jsonRequestDecoder(cx, &article)

	if err := c.UpdateId(id, article); err != nil {
		log.Println("Error: %s", err.Error())
		return goweb.Respond.WithStatus(cx, http.StatusForbidden)
	}

	log.Printf("Updated a article id=%s", id)
	return goweb.Respond.WithOK(cx)
}

func (cr *Controller) UpdateMany(cx context.Context) error {
	log.Println("Update all articles...")
	c := cr.db.C(COLLECTION)

	var articles []*Article
	jsonRequestDecoder(cx, &articles)

	if _, err := c.UpdateAll(nil, articles); err != nil {
		log.Println("Error: %s", err.Error())
		return goweb.Respond.WithStatus(cx, http.StatusForbidden)
	}

	log.Println("Updated all articles")
	return goweb.Respond.WithOK(cx)
}
