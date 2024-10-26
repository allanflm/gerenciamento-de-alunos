package routes

import (
	"github.com/allanflm/api-go-gim/controllers"
	"github.com/gin-gonic/gin"
)

func HangleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.Run()
}
