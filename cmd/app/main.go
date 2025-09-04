package main

import (
	"fmt"

	"github.com/Rocksus/fundtract/internal/platform/storage/sqlite"
	"github.com/labstack/echo/v4"
)

//go:generate sqlc generate -f ../../sqlc.yaml
//go:generate swag init --pd -g app.go -o ../../docs -d ./,../../internal
//go:generate swag fmt

func InitApp() *echo.Echo {
	e := echo.New()

	_ = sqlite.New()

	return e
}

func main() {
	app := InitApp()
	app.Logger.Fatal(app.Start(fmt.Sprint(":", 14045)))
}
