package main

import (
	"fmt"

	"github.com/JueViGrace/bakery-server/internal/server"
)

func main() {
	server := server.New()

	if err := server.Init(); err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
