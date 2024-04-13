Alterar os caminhos de import de acordo com o seu pc:

Exemplo:

"github.com/hugohvm25/API-GO-GIN/database"
"github.com/hugohvm25/API-GO-GIN/routes"


=== DOCKER ===

Para rodar o projeto corretamente com o docker e não ter incompatibilidade da versão do BD, excluir a pasta postgres-data e executar os códigos:

docker-compose build
ocker-compose up 


=== instalando o pacote Validador ===

https://pkg.go.dev/gopkg.in/validator.v2

executar o comando: go get gopkg.in/validator.v2