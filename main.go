package main

import "golang-demo/initialize"

func main()  {
	r := initialize.Routers()
	r.Run(":8080")
}



