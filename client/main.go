package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Cotacao struct {
    Cambio string `json:"cambio"`
}

func main() {
	c := http.Client{}

	req, err := http.NewRequest("GET", "http://localhost:8080/cotacao", nil)

	if err != nil {
		panic(err)
	}

	resp, err := c.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))

	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)

	if err != nil {
		panic(err)
	}

	file, err := os.Create("cotacao.txt")
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
	}

	defer file.Close()
	
	fmt.Print("Arquivo criado com sucesso!")

	_, err = file.WriteString(fmt.Sprintf("Dólar: %v", cotacao.Cambio))
	if err != nil {
		panic(err)
	}

	fmt.Print("Dados gravados com sucesso!")
}