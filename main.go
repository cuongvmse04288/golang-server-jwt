package main

import (
	"golang-demo/initialize"
	"log"
)

func main() {
	r := initialize.Routers()
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
