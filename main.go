package main

import (
	"net/http"
	"store/routes"
)

func main() {
	routes.LoadRouters()
	http.ListenAndServe(":8000", nil)
}
