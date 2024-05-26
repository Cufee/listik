package main

import (
	"os"

	"github.com/cufee/shopping-list/internal/logic"
	"github.com/cufee/shopping-list/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

//go:generate go run github.com/steebchen/prisma-client-go generate

// // go:embed assets
// var assets embed.FS

func main() {
	client, err := logic.NewDatabaseClient()
	if err != nil {
		panic(err)
	}

	// s := server.New(client, echo.MustSubFS(assets, "assets"))
	s := server.New(client, nil)
	panic(s.Start(":" + os.Getenv("PORT")))
}
