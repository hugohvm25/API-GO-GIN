package main

import (
	"github.com/hugohvm25/API-GO-GIN/database"
	"github.com/hugohvm25/API-GO-GIN/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
