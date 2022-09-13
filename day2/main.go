package main

import (
	"github.com/arioki1/alterra-agmc/config"
	"github.com/arioki1/alterra-agmc/routes"
)

func init() {
	config.InitDB()
}
func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
