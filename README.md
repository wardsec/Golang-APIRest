
# RESTApi - Golang

Este projeto foi realizado para testar e consolidar conhecimenso em RESTApis em Go.
Com este projeto você é capaz de Criar, editar e deletar Dados de ruas com nomes de 'personalidades'
obviamente você pode editar o codigo da maneira que desejar.

## Como utilizar
1 - Temos dois arquivos, 1 sendo nosso frontend em React e Nosso backend em Golang.
Para iniciar nosso frontend 
## Rodando o projeto

Antes de tudo, Crie a estrutura de banco da dados 
```
  create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades(nome, historia) VALUES
('Deodato Petit Wertheimer', 'Deodato Petit Wertheimer foi um médico e político brasileiro, seus primeiros anos de vida foram em São Paulo, mas logo mudou para Nova Friburgo no Estado do Rio de Janeiro e com 11 anos de idade ingressou no Colégio Anchieta dos jesuítas.'),
('Carmela Dutra', 'Carmela Teles Leite Dutra foi a primeira-dama do Brasil, de 31 de janeiro de 1946 até a sua morte, tendo sido a esposa de Eurico Gaspar Dutra, 16.º Presidente do Brasil. Era, carinhosamente, chamada de Dona Santinha, pela sua forte religiosidade, fazendo seu marido abrir uma capelinha no Palácio Guanabara.');
```

Clone o projeto

```bash
  git clone https://link-para-o-projeto
```

Inciando o front end - Apos entrar na pasta inicia o rode o comando:

```bash
  npm install
  npm start

```
Apos isso entre na pasta de nosso back end e configure a conexao com o banco de dados em database/db.go
```bash
  go run main.go
```
Rode o comando acima para inicar o back end



## Funcionalidades

- Criar, deletar e fazed update.


## Screenshots

![App Screenshot](<img src="https://i.ibb.co/K7PpSH1/Web-capture-14-3-2022-174251-localhost.jpg")

