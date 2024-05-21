/*

a função de teste do GO deve sempre ter a palavra Test com T maiúsculo na primeira letra

== FUNÇÃO DE EXEPLO DE TESTE ==
func TestFalhador(t *testing.T) {
	t.Fatalf("Teste falhou e propósito, não se preocupe!")
}

PARA EXECUTAR O TESTE SOMENTE DO QUE DESEJA

go test -run TestVerificaçãoDaSaudacaoComParametro

*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hugohvm25/API-GO-GIN/controllers"
	"github.com/hugohvm25/API-GO-GIN/database"
	"github.com/hugohvm25/API-GO-GIN/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotasTeste() *gin.Engine {
	//para melhorar a visibilidade da resposta do teste é possivel usar o Modo Release de forma compacta
	gin.SetMode(gin.ReleaseMode)
	// rota sem cadatro padrão
	rotas := gin.Default()
	return rotas
}

func TestVerificaçãoDaSaudacaoComParametro(t *testing.T) {
	//cria uma nova rota a partir da função base de SETUP de Rotas
	r := SetupRotasTeste()
	r.GET("/:nome", controllers.Saudacao)
	//requisição (tipo de requisição, end point da requisição, algum tipo de dado que queira passar (json, etc.))
	req, _ := http.NewRequest("GET", "/hugo", nil)
	//vai implementar a interface de quem vai armazenar a resposta
	resposta := httptest.NewRecorder()
	//requisição parâmetros(onde guardar a resposta da requisição, qual tipo de requisição)
	r.ServeHTTP(resposta, req)
	//a partir do testify, usar o assert para facilitar e encurtar o código deixando mais limpo
	assert.Equal(t, http.StatusOK, resposta.Code, "Os retornos deveriam ser iguais!")
	//retorno com a informação para verificação - fazer encapsulamento do testo com crase ``
	mockDaResposta := `{"API diz:":"E ai hugo, tudo beleza?"}`
	//retorno com a leitura da resposta
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	//convertendo para string para que ele não retorne a resposta em bites
	assert.Equal(t, mockDaResposta, string(respostaBody))
	fmt.Println("Mensagem esperada:", mockDaResposta)
	fmt.Println("Mensagem apresentada:", string(respostaBody))

	// if resposta.Code != http.StatusOK {
	// 	t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	// } else {
	// 	fmt.Println("Passou no teste")
	// }
}

func CriaAlunoMock() {
	//passando os dados para armazenamento no banco de dados
	aluno := models.Aluno{Nome: "Nome: TesteAluno", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	//como foi declarado publicamente antes de todos os códigos não é necessário o :
	ID = int(aluno.ID)
}

func DeletarAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)

}
func TestListandoTodosAlunosHandler(t *testing.T) {
	//conexão com o banco de dados da aplicação
	database.ConectaComBancoDeDados()
	//cria o aluno no banco a partir da função
	CriaAlunoMock()
	//deletar depois que rodar esta função
	defer DeletarAlunoMock()
	//cria a rota de teste
	r := SetupRotasTeste()
	//busca o caminho da requisição para realizar o teste
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	//requisição = metodo da requisição, caminho (path) a ser utilizado, e mensagem (conteúdo de retorno)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	//para armazenar o retorno da resposta
	resposta := httptest.NewRecorder()
	//requisição
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	//print para confirmar se está buscando a informação correta no banco de dados
	fmt.Println(resposta.Body)
}

/*
preciso de quais parametros para função de teste:
- preciso de conexão com banco de dados?
- preciso criar algum dado?
- preciso deletar após executar a função?
*/

func BuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletarAlunoMock()
	r := SetupRotasTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/aluno/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
