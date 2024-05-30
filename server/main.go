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
        s.CotacaoHandler(w, r)
	default:
        http.NotFound(w, r)
	}
}

func (s server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte(s.welcomeMessage))
}

func (s server) CotacaoHandler(w http.ResponseWriter, r *http.Request) {
    ultimaCotacao, err := BuscaCotacao()
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
    }

    resposta := RespostaCotacao{
        Cambio: ultimaCotacao.USDBRL.Bid,
    }

    w.Header().Set("Content-Type", "application/json")
	  w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(resposta)
}

func BuscaCotacao() (*Cotacao, error) {
    resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var cotacao Cotacao
    err = json.Unmarshal(body, &cotacao)
    if err != nil {
        return nil, err
    }

    return &cotacao, nil
}
