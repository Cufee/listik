package main

import (
	"os"

	"github.com/cufee/shopping-list/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	s := server.New()
	s.Start(":" + os.Getenv("PORT"))
}
