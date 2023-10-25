package routes

import (
	"api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ListaAlunos)
	r.POST("/alunos", controllers.CriaAluno)
	r.Run()
}
