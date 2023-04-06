package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {

	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.") //executa quando o timeout expira e o cancel() for executado
		return
	case <-time.After(3 * time.Second): //Aguarda 2 segundos. Se o time for maior que o timeout o contexto é cancelado
		fmt.Println("Hotel booked.")
	}
}
