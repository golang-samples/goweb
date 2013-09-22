package main

import (
	"fmt"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"net/http"
)

type MyController struct{}

func (cr *MyController) Create(cx context.Context) error {
	body := "Create resource"
	return goweb.Respond.With(cx, http.StatusCreated, []byte(body))
}

func (cr *MyController) Read(id string, cx context.Context) error {
	body := fmt.Sprintf("Read resource %s", id)
	return goweb.Respond.With(cx, http.StatusOK, []byte(body))
}

func (cr *MyController) ReadMany(cx context.Context) error {
	body := "Read many resources"
	return goweb.Respond.With(cx, http.StatusOK, []byte(body))
}

func (cr *MyController) Update(id string, cx context.Context) error {
	body := fmt.Sprintf("Update resource %s", id)
	return goweb.Respond.With(cx, http.StatusOK, []byte(body))
}

func (cr *MyController) UpdateMany(cx context.Context) error {
	body := "Update many resources"
	return goweb.Respond.With(cx, http.StatusOK, []byte(body))
}

func (cr *MyController) Delete(id string, cx context.Context) error {
	body := fmt.Sprintf("Delete resource %s", id)
	return goweb.Respond.With(cx, http.StatusOK, []byte(body))
}

func (cr *MyController) DeleteMany(cx context.Context) error {
	body := "Delete many resources"
	return goweb.Respond.With(cx, http.StatusOK, []byte(body))
}

func main() {
	controller := new(MyController)
	goweb.MapController("/api", controller)
	http.ListenAndServe(":8080", goweb.DefaultHttpHandler())
}
