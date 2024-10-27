package main

import (
	"github.com/allanflm/api-go-gim/database"
	"github.com/allanflm/api-go-gim/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HangleRequests()
}
