package main

import (
	"fmt"

	"github.com/JueViGrace/bakery-go/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server := server.New()

	err := server.Init()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
