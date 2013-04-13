package main

import (
	"fmt"
	"github.com/stretchrcom/goweb/goweb"
)

type MyController struct{}

func (cr *MyController) Create(cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Create resource")
}

func (cr *MyController) Read(id string, cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Read resource %s", id)
}

func (cr *MyController) ReadMany(cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Read many resources")
}

func (cr *MyController) Update(id string, cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Update resource %s", id)
}

func (cr *MyController) UpdateMany(cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Update many resources")
}

func (cr *MyController) Delete(id string, cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Delete resource %s", id)
}

func (cr *MyController) DeleteMany(cx *goweb.Context) {
	fmt.Fprintf(cx.ResponseWriter, "Delete many resources")
}

func main() {
	controller := new(MyController)
	goweb.MapRest("/api", controller)
	goweb.ListenAndServe(":8080")
}
