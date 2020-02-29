package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(55 * time.Microsecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("game over")
	case <-time.After(1 * time.Second):
		fmt.Println("you question error")
	}
}
