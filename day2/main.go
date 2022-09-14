package main

import (
	"github.com/arioki1/alterra-agmc/day2/config"
	"github.com/arioki1/alterra-agmc/day2/routes"
)

func init() {
	config.InitDB()
}
func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
