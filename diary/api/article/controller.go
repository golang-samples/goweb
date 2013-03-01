package article

import (
	"github.com/stretchrcom/goweb/goweb"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

type Controller struct {
	db *mgo.Database
}

func NewController(db *mgo.Database) *Controller {
	if db == nil {
		panic("database cannot be nil")
	}

	return &Controller{db}
}

func (cr *Controller) Create(cx *goweb.Context) {
	log.Println("Creating a article...")
	c := cr.db.C(COLLECTION)

	var article Article
	decoder := new(goweb.JsonRequestDecoder)
	decoder.Unmarshal(cx, &article)

	article.Id = bson.NewObjectId().Hex()

	if err := c.Insert(&article); err != nil {
		log.Println("Error: %s", err.Error())
		cx.RespondWithError(http.StatusForbidden)
		return
	}

	log.Printf("Created diary id=%s", article.Id)
	cx.RespondWithData(article)
}

func (cr *Controller) Delete(id string, cx *goweb.Context) {
	log.Printf("Deleting a article id=%s...", id)
	c := cr.db.C(COLLECTION)
	if err := c.RemoveId(id); err != nil {
		log.Printf("Error: %s", err.Error())
		cx.RespondWithError(http.StatusForbidden)
		return
	}

	log.Printf("Deleted article id=%s", id)
	cx.RespondWithOK()
}

func (cr *Controller) DeleteMany(cx *goweb.Context) {
	log.Println("Deleting all articles...")
	c := cr.db.C(COLLECTION)
	if _, err := c.RemoveAll(nil); err != nil {
		log.Println("Error: %s", err.Error())
		cx.RespondWithError(http.StatusForbidden)
		return
	}

	log.Println("Deleted all articles")
	cx.RespondWithOK()
}

func (cr *Controller) Read(id string, cx *goweb.Context) {
	log.Printf("Read a article id=%s", id)
	c := cr.db.C(COLLECTION)
	var article Article
	if err := c.FindId(id).One(&article); err != nil {
		log.Println("Error: %s", err.Error())
		cx.RespondWithError(http.StatusForbidden)
		return
	}

	log.Printf("Read article id=%s", id)
	cx.RespondWithData(article)
}

func (cr *Controller) ReadMany(cx *goweb.Context) {
	log.Println("Read all articles...")
	c := cr.db.C(COLLECTION)
	count, err := c.Count()
	if err != nil {
		log.Println("Error: %s", err.Error())
		cx.RespondWithError(http.StatusForbidden)
		return
	}

	articles := make([]*Article, count)
	if err := c.Find(nil).All(&articles); err != nil {
		log.Println("Error: %s", err.Error())
		cx.RespondWithError(http.StatusForbidden)
		return
	}

	log.Printf("Read all %d articles", count)
	cx.RespondWithData(articles)
}

func (cr *Controller) Update(id string, cx *goweb.Context) {
	log.Printf("Update a article id=%s...", id)
	c := cr.db.C(COLLECTION)

	var article *Article
	decoder := new(goweb.JsonRequestDecoder)
	decoder.Unmarshal(cx, &article)

	if err := c.UpdateId(id, article); err != nil {
		log.Println("Error: %s", err.Error())
		cx.RespondWithError(http.StatusForbidden)
		return
	}

	log.Printf("Updated a article id=%s", id)
	cx.RespondWithOK()
}

func (cr *Controller) UpdateMany(cx *goweb.Context) {
	log.Println("Update all articles...")
	c := cr.db.C(COLLECTION)

	var articles []*Article
	decoder := new(goweb.JsonRequestDecoder)
	decoder.Unmarshal(cx, &articles)

	if _, err := c.UpdateAll(nil, articles); err != nil {
		log.Println("Error: %s", err.Error())
		cx.RespondWithError(http.StatusForbidden)
		return
	}

	log.Println("Updated all articles")
	cx.RespondWithOK()
}
