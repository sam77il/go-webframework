package main

import (
	"sugarweb.dev/framework/pages"
	"sugarweb.dev/framework/sugar"
)

func main() {
	server := sugar.Init()

	server.Routes.GET("/", handleRoot)
	server.Routes.GET("/products", handleProductsPage)

	server.Routes.GET("/favicon.ico", func(r *sugar.RouteHandler) {
    	r.Status(204)
	})
	server.Listen(":8080")
}

func handleRoot(r *sugar.RouteHandler) {
	r.Status(200).HTML(pages.RootPage())
}

func handleProductsPage(r *sugar.RouteHandler) {
	r.Status(200).HTML(pages.ProductsPage())
}