package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"wfw"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%2d-%2d", year, month, day)
}

func main() {
	r := wfw.New()

	r.Use(wfw.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Tom", Age: 20}
	stu2 := &student{Name: "Amy", Age: 18}
	r.GET("/", func(c *wfw.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *wfw.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", wfw.H{
			"title":  "wfw",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.GET("/date", func(c *wfw.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", wfw.H{
			"title": "wfw",
			"now":   time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		})
	})

	log.Fatal(r.Run(":9999"))
}
