package main

import (
	"fmt"

	mongoinit "ProductService/Mongo"
	router "ProductService/Router"
)

func init() {
	mongoinit.MongoConnInit()
}

func StartServer() {
	routes := router.RouteDispatcher()
	fmt.Println("start")
	routes.Run(":9090")
}

func main() {
	StartServer()
}
