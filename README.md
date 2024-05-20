Alterar os caminhos de import de acordo com o seu pc:

Exemplo:

"github.com/hugohvm25/API-GO-GIN/database"
"github.com/hugohvm25/API-GO-GIN/routes"


=== DOCKER ===

Para rodar o projeto corretamente com o docker e não ter incompatibilidade da versão do BD, excluir a pasta postgres-data e executar os códigos:

docker-compose build
docker-compose up 


=== instalando o pacote Validador ===

https://pkg.go.dev/gopkg.in/validator.v2

executar o comando: go get gopkg.in/validator.v2

=== criando arquivos de testes ===

go test

a função de teste do GO deve sempre ter a palavra Test com T maiúsculo na primeira letra
func TestNome_do_teste(t *testing.T)


biblioteca para instalação que possui alguns tipos de parametros já pré definidos
go get github.com/stretchr/testify 
