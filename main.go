package main

import (
	"deploy/route"
)

func main() {
	route := route.StartRoute()
	route.Start(":8000")
}
