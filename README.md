# Go Project

Este repositório tem como objetivo demonstrar **noções de arquitetura**, **manipulação de banco de dados** e **testes unitários** utilizando a linguagem **Go**.  

## Objetivo do Projeto

- Aplicar conceitos de **arquitetura de software** em Go.  
- Demonstrar como interagir com **bancos de dados** de forma organizada.  
- Implementar **testes unitários** para garantir confiabilidade do código.  
- Futuramente, integrar:
  - **Filas e processamento assíncrono**  
  - **Redis** para persistência de cache e filas  
  - **Kubernetes** para orquestração de containers  
  - **CI/CD** com **CircleCI** e **Jenkins**

## 🛠 Tecnologias
- Banco de dados relacional (PostgreSQL, MySQL ou outro à sua escolha)  
- Testes unitários com `testing` do Go  
- Futuras integrações: Redis, Kubernetes, CircleCI, Jenkins

## Estrutura do Projeto

```text
.
├── cmd/
│   └── server/
│       └── config.go           # Configurações da aplicação
├── domain/
│   └── user.go                 # Modelos/Domínios
├── external/
│   └── aws/
│       └── s3.go               # Integração com AWS S3
├── http/
│   ├── controller/
│   │   ├── check.go            # Controller de verificação da saúde da aplicação
│   │   └── user.go             # Controller que gerencia as regras de negócios dos usuários
│   ├── handler/
│   │   └── handler.go          # Lista de rotas
│   ├── middleware/
│   │   └── middleware.go       # Middlewares
│   └── router/
│       └── router.go           # Configura e retorna o roteador HTTP do Gin com CORS 
├── service/
│   ├── user.go                 # Lógica de negócio
│   └── user_test.go            # Testes unitários da service de usuários
├── storage/
│   ├── database/
│   │   └── postgresql.go       # Conexão com o banco de dados
│   └── repository/
│       └── user.go             # Repositório de usuários
├── .gitignore
├── cover.txt
├── Dockerfile
├── go.mod
├── go.sum
└── main.go

```

## Como rodar

Clone o repositório:

```bash
git clone https://github.com/Dyckson/go-projects.git
cd go-projects
```

Instale as dependências e execute a aplicação:

```bash
go get
go run main.go
```

## Testes

Para rodar os testes unitários:

```bash
go test ./... -v
```
