package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()                               // contexto vazio
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond) // se passa o tempo do timeout o contexto é cancelado antes de chamar o cancel()
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req) // cria o objeto client
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
