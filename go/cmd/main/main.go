package main

import (
	_ "go/config"
	"go/pkg/router"
	"os"
)

func main() {
	e := router.New()

	e.Logger.Fatal(e.Start(os.Getenv("SERVER_PORT")))
}
