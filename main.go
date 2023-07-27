package main

import (
	"log"
	"net/http"
	"wfw"
)

func main() {
	engine := new(wfw.Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
