# CI_EC2 🚀
![Badge da linguagem Go](https://img.shields.io/badge/Linguagem-Go-yellowgreen)

## Descrição 📝
Este projeto foi criado para demonstrar uma infraestrutura de Continuous Integration (CI) utilizando o serviço de Amazon EC2. O objetivo é fornecer uma experiência de desenvolvimento e deployment simplificada e escalável para aplicativos modernos.

## Funcionalidades Principais 🤖
O projeto inclui as seguintes funcionalidades:

* Deploy em ambiente Dockerizado utilizando o Compose.
* Integração contínua utilizando o GitHub Actions.
* Conexão com banco de dados PostgreSQL.
* Utilização da linguagem Go com o framework Gin.

## Pré-requisitos 📋
Para rodar o projeto, você precisará ter:

* Docker instalado no computador.
* Go instalado no computador.
* Conta no GitHub para configurar o CI.
* Conta no AWS para criar uma instancia EC2.

## Instalação 🔧
Para instalar o projeto, siga os passos abaixo:

1. Clonar o repositório em seu computador utilizando o comando `git clone`.
2. Criar um arquivo `.env` com as variáveis de ambiente necessárias.
3. Criar uma instancia EC2 no AWS.
4. Configurar o CI no GitHub Actions.
5. Rodar o comando `docker-compose up -d` para iniciar o container.
6. Acessar o endereço `localhost:8000` no navegador para visualizar o aplicativo.

## Exemplos de uso 📈
O projeto inclui uma API RESTful para exemplo. Você pode consultar o endpoint `/users` utilizando o método GET.

## Estrutura do projeto 🗂
A estrutura do projeto é a seguinte:
```
CI_EC2
├── .github
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── main.go
```

## Licença 📜
Este projeto é licenciado sob a Licença MIT. O texto da licença pode ser encontrado no arquivo `LICENSE`.

## Agradecimentos 🙏
Obrigado pela oportunidade de participar do projeto.