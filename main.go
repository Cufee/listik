package main

import (
	"embed"
	"os"

	"github.com/cufee/shopping-list/internal/server"
	"github.com/labstack/echo/v4"

	_ "github.com/joho/godotenv/autoload"
)

//go:embed assets
var assets embed.FS

func main() {
	fs := echo.MustSubFS(assets, "assets")
	s := server.New(fs)
	s.Start(":" + os.Getenv("PORT"))
}
