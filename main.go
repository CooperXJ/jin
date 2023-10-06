package main

import (
	"fmt"
	"net/http"
)

func main() {
	e := Default()
	e.LoadHTMLGlob("test/*")

	e.Use(Logger())

	e.GET("/hello", func(c *Context) {
		fmt.Fprintf(c.Writer, "success\n")
	})

	v1 := e.Group("/v1")
	v1.GET("/", func(c *Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	v1.GET("/hello", func(c *Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	e.Static("/assets", "./test/")
	e.Run(":9999")
}
