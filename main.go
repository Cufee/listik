package main

import (
	"os"

	"github.com/cufee/shopping-list/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

//go:generate go run github.com/steebchen/prisma-client-go generate

func main() {
	s := server.New(nil)
	s.Start(":" + os.Getenv("PORT"))
}
