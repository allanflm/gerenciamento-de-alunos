# Projeto de Gerenciamento de Alunos em Go

Este projeto é uma API de gerenciamento de alunos desenvolvida em Go, usando o framework Gin-Gonic para criação de endpoints REST e o GORM para manipulação de um banco de dados PostgreSQL. A aplicação permite o cadastro, consulta, atualização e exclusão de dados de alunos. Toda a infraestrutura do projeto é gerenciada com Docker e Docker Compose, facilitando a configuração do ambiente.

## Tecnologias Utilizadas

- **Golang**: Linguagem principal do projeto.
- **Gin-Gonic**: Framework para criação de APIs HTTP.
- **GORM**: ORM para manipulação de dados com suporte ao PostgreSQL.
- **PostgreSQL**: Banco de dados utilizado para persistência dos dados.
- **PgAdmin**: Interface gráfica para gerenciar o banco de dados.
- **Docker**: Containerização do banco de dados e da aplicação.
- **Docker Compose**: Gerenciamento da infraestrutura do projeto.

## Funcionalidades

1. **Listar Alunos**: Exibe todos os alunos cadastrados.
2. **Cadastrar Novo Aluno**: Permite o cadastro de um novo aluno.
3. **Consultar Aluno por ID**: Retorna os dados de um aluno específico pelo seu ID.
4. **Consultar Aluno por CPF**: Retorna os dados de um aluno específico pelo CPF.
5. **Atualizar Aluno**: Edita os dados de um aluno existente.
6. **Excluir Aluno**: Remove um aluno do sistema.

## Configuração

### Pré-requisitos

- Docker
- Docker Compose
- Golang 1.16+

### Passos para Executar

1. **Clone o Repositório**:
   ```bash
   git clone https://github.com/allanflm/gerenciamento-de-alunos.git
   cd gerenciamento-de-alunos
2. **Inicie os Contêineres**:
O arquivo docker-compose.yml gerencia os contêineres para PostgreSQL e PgAdmin. Para iniciar o ambiente:

``bash
Copiar código
docker-compose up -d
Execute o Projeto:
Com o banco de dados em execução, execute o comando:

bash
Copiar código
go run main.go
Variáveis de Ambiente
Atualize a string de conexão com o banco de dados conforme sua necessidade no arquivo de configuração database/database.go.

Endpoints
Método	Endpoint	Descrição
GET	/alunos	Lista todos os alunos
GET	/:nome	Saudação personalizada
POST	/alunos	Cria um novo aluno
GET	/alunos/:id	Busca aluno pelo ID
DELETE	/alunos/:id	Deleta um aluno
PATCH	/alunos/:id	Edita dados de um aluno
GET	/alunos/cpf/:cpf	Busca aluno pelo CPF
