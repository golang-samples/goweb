package main

import (
	"fmt"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"net/http"
)

func main() {
	handler := func(c context.Context) error {
		name := c.PathParams().Get("name")
		animal := c.PathParams().Get("animal")
		body := fmt.Sprintf("Hey %s, your favorite animal is a %s", name, animal)
		return goweb.Respond.With(c, http.StatusOK, []byte(body))
	}

	goweb.Map("/people/{name}/animals/{animal}", handler)
	http.ListenAndServe(":8080", goweb.DefaultHttpHandler())
}
