package controllers

import (
	"api-gin/models"

	"github.com/gin-gonic/gin"
)

func ListaAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Olá: ": nome,
	})
}
