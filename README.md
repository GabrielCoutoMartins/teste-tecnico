## ▶️ Como executar o projeto


- Go 1.25
- Banco de dados postgres
- javascript
- instalar a lib de conexao com o postgres e de migration
-go mod tidy
-go get github.com/lib/pq
-go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest 
- Criar um banco de dados no postgres
- Configure a string de conexão em db/db.go
- utilizar o comando da pasta migration alterando a string de conexão tambem
-migrate -path ./migrations -database "postgres://postgres:gabriel123@localhost:5432/tasks?sslmode=disable" up
- utilizar o comando para rodar a aplicação
-go run main.go
- http://localhost:8080
