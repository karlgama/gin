package main

import (
	"api-gin/models"
	"api-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{
			Nome: "teste",
			CPF:  "234",
			RG:   "45",
		},
		{
			Nome: "teste2",
			CPF:  "234",
			RG:   "45",
		},
	}
	routes.HandleRequests()
}
