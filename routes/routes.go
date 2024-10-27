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

	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)

	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	r.PATCH("/alunos/:id", controllers.EditarAluno)

	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run()
}
