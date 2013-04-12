package main

import (
	"fmt"
	"github.com/stretchrcom/goweb/goweb"
)

func main() {

	handler := func(c *goweb.Context) {
		name := c.PathParams["name"]
		animal := c.PathParams["animal"]
		fmt.Fprintf(c.ResponseWriter, "Hey %s, your favorite animal is a %s", name, animal)
	}

	goweb.MapFunc("/people/{name}/animals/{animal}", handler)
	goweb.ListenAndServe(":8080")
}
