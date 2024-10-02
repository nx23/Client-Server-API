
# Desafio Golang - Webserver com Cotação do Dólar

Este projeto foi desenvolvido em Golang para demonstrar o uso de web servers HTTP, contextos, manipulação de arquivos e integração com APIs externas, além de persistência de dados em banco de dados. A aplicação é dividida em dois sistemas: `client.go` e `server.go`.

## Funcionalidade

### server.go

O `server.go` é responsável por:
- Consumir a API pública de câmbio [AwesomeAPI](https://economia.awesomeapi.com.br/json/last/USD-BRL) para obter a cotação do dólar em relação ao real (USD/BRL).
- Registrar no banco de dados SQLite cada cotação recebida, com um timeout máximo de 10ms para a inserção.
- A comunicação com a API de câmbio tem um timeout de 200ms, utilizando o pacote `context`.
- O endpoint `/cotacao` na porta `8080` retorna a cotação atual (campo `bid` do JSON retornado pela API) no formato JSON.
- Em caso de falhas de timeout, o server loga erros, indicando o tempo insuficiente para concluir a operação.

### client.go

O `client.go` realiza uma requisição HTTP para o `server.go` buscando a cotação do dólar:
- O cliente se conecta ao servidor na rota `/cotacao` e recebe a cotação no formato JSON.
- Utiliza o campo `bid` da resposta JSON para salvar o valor da cotação em um arquivo de texto `cotacao.txt` no formato: `Dólar: {valor}`.
- O cliente tem um timeout máximo de 300ms para receber o valor da cotação do servidor, também utilizando o pacote `context`.
- Caso ocorra um timeout, o erro é registrado nos logs.

## Como executar a aplicação

### Requisitos
- Go 1.16+
- SQLite3

### Passo a passo

1. Clone o repositório:
    ```bash
    git clone https://github.com/nx23/Client-Server-API
    cd Client-Server-API
    ```

2. Instale as dependências necessárias:
    ```bash
    go mod tidy
    ```

3. Execute o servidor:
    ```bash
    go run server/main.go
    ```

4. Em outro terminal, execute o cliente:
    ```bash
    go run client/main.go
    ```

## Estrutura do Projeto

- `client.go`: Código do cliente responsável por fazer a requisição ao servidor e salvar a cotação em um arquivo.
- `server.go`: Código do servidor responsável por consumir a API de câmbio, persistir as informações no banco de dados e fornecer a cotação via HTTP.
- `cotacao.txt`: Arquivo gerado pelo cliente contendo a cotação do dólar.
- `cotacoes.db`: Banco de dados SQLite utilizado pelo servidor para armazenar as cotações.

## Tecnologias Utilizadas

- **Golang**: Linguagem principal para o desenvolvimento.
- **HTTP**: Para comunicação entre cliente e servidor.
- **SQLite**: Banco de dados leve para persistir as cotações.
- **Context**: Para gestão de timeouts e cancelamento de operações.
- **API [AwesomeAPI](https://docs.awesomeapi.com.br/api-de-moedas)**: Fonte dos dados de cotação.

## Logs e Erros

A aplicação registra logs quando ocorre algum erro de timeout:
- **Timeout ao consumir a API de câmbio**: Caso a resposta demore mais de 200ms.
- **Timeout ao persistir no banco de dados**: Caso a inserção no SQLite demore mais de 10ms.
- **Timeout no cliente**: Caso o cliente não receba uma resposta do servidor em até 300ms.

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

---

## Licença

Este projeto está sob a licença MIT.

---
