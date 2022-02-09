package main

import (
	"gvblog/model"
	"gvblog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
