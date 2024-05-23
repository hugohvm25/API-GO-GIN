package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugohvm25/API-GO-GIN/controllers"
)

func HandleRequests() {
	r := gin.Default()
	//pasta onde está a página html
	r.LoadHTMLGlob("templates/*")
	//configando o GIN para arquivos estaticos
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	//rota de exibição da pagina html
	r.GET("/index", controllers.ExibePaginaIndex)
	//configuração para rotas não encontradas
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run()
}
