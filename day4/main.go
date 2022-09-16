package main

import (
	"github.com/arioki1/alterra-agmc/day4/config"
	m "github.com/arioki1/alterra-agmc/day4/middlewares"
	"github.com/arioki1/alterra-agmc/day4/routes"
)

func init() {
	config.InitDB()
}

func main() {
	e := routes.New()
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
