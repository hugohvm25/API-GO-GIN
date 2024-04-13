/*

a função de teste do GO deve sempre ter a palavra Test com T maiúsculo na primeira letra

== FUNÇÃO DE EXEPLO DE TESTE ==
func TestFalhador(t *testing.T) {
	t.Fatalf("Teste falhou e propósito, não se preocupe!")
}

*/

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hugohvm25/API-GO-GIN/controllers"
)

func SetupRotasTeste() *gin.Engine {

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
	if resposta.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	} else {
		fmt.Println("Passou no teste")
	}

}
