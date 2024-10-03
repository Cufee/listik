package main

import (
	"embed"
	"os"

	"github.com/cufee/shopping-list/internal/logic"
	"github.com/cufee/shopping-list/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

//go:generate go run github.com/steebchen/prisma-client-go generate
//go:generate templ generate

// go:embed assets/*
var assets embed.FS

func main() {
	client, err := logic.NewDatabaseClient()
	if err != nil {
		panic(err)
	}

	s := server.New(client, assets)
	panic(s.Start(":" + os.Getenv("PORT")))
}
