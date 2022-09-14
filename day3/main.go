package main

import (
	"github.com/arioki1/alterra-agmc/day3/config"
	m "github.com/arioki1/alterra-agmc/day3/middlewares"
	"github.com/arioki1/alterra-agmc/day3/routes"
)

func init() {
	config.InitDB()
}

func main() {
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
