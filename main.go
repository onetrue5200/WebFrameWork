package main

import (
	"log"
	"wfw"
)

func main() {
	r := wfw.New()

	r.Static("/assets", "./static")

	log.Fatal(r.Run(":9999"))
}
