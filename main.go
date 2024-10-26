package main

import (
	"github.com/allanflm/api-go-gim/database"
	"github.com/allanflm/api-go-gim/models"
	"github.com/allanflm/api-go-gim/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Allan", CPF: "12345678921", RG: "00000000"},
		{Nome: "Felipe", CPF: "000 000 000 00 ", RG: "00000021"},
	}
	routes.HangleRequests()
}
