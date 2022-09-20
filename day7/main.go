package main

import (
	"fmt"
	"github.com/arioki1/alterra-agmc/day7/config"
	m "github.com/arioki1/alterra-agmc/day7/middlewares"
	"github.com/arioki1/alterra-agmc/day7/routes"
	"os"
)

func init() {
	config.InitDB()
}

func main() {
	e := routes.New()
	m.LogMiddlewares(e)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
