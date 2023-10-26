package routes

import (
	"api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ListaAlunos)
	r.GET("/alunos/:id", controllers.BuscaPorId)
	r.POST("/alunos", controllers.CriaAluno)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	r.PUT("/alunos/:id", controllers.EditaAluno)
	r.Run()
}
