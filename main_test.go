/*

a função de teste do GO deve sempre ter a palavra Test com T maiúsculo na primeira letra

*/

package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupRotasTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestFalhador(t *testing.T) {
	t.Fatalf("Teste falhou e propósito, não se preocupe!")

}
