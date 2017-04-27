package main

import (
	"net/http"
	"html/template"
	_ "fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main()  {
	router := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	t := template.Must(template.New("templates").Delims("${{", "}}").ParseGlob("resources/templates/*.tmpl.html"))
	router.SetHTMLTemplate(t)

	router.Static("static/js", "./resources/js")

	pageGet(router, "index")

	router.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})

	router.Run(":8082")
}

func pageGet(router* gin.Engine, name string)  {
	router.GET(name, func(c* gin.Context) {
		session := sessions.Default(c)

		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()

		fmt.Printf("count %d\n", count)

		c.HTML(http.StatusOK, name + ".tmpl.html", gin.H{})
	})
}