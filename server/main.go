package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", server{welcomeMessage: "Bem vindo a minha API de Cotacao"})
	http.ListenAndServe(":8080", mux)
}

type Cotacao struct {
    USDBRL struct {
        Code       string `json:"code"`
        Codein     string `json:"codein"`
        Name       string `json:"name"`
        High       string `json:"high"`
        Low        string `json:"low"`
        VarBid     string `json:"varBid"`
        PctChange  string `json:"pctChange"`
        Bid        string `json:"bid"`
        Ask        string `json:"ask"`
        Timestamp  string `json:"timestamp"`
        CreateDate string `json:"create_date"`
    } `json:"USDBRL"`
}

type RespostaCotacao struct {
    Cambio string `json:"cambio"`
}

type server struct {
	welcomeMessage string
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
        s.HomeHandler(w, r)
	case "/cotacao":
        s.Cotacao(w, r)
	default:
        http.NotFound(w, r)
	}
}

func (s server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte(s.welcomeMessage))
}

func (s server) Cotacao(w http.ResponseWriter, r *http.Request) {
    ultimaCotacao, error := BuscaCotacao()
    if error != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", error)
    }

    resposta := RespostaCotacao{
        Cambio: ultimaCotacao.USDBRL.Bid,
    }

    w.Header().Set("Content-Type", "application/json")
	  w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(resposta)
}

func BuscaCotacao() (*Cotacao, error) {
    resp, error := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
    if error != nil {
        return nil, error
    }
    defer resp.Body.Close()

    body, error := io.ReadAll(resp.Body)
    if error != nil {
        return nil, error
    }

    var cotacao Cotacao
    error = json.Unmarshal(body, &cotacao)
    if error != nil {
        return nil, error
    }

    return &cotacao, nil
}
