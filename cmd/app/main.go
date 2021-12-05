package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lucasmmo/setup-go/pkg/app"
)

var (
	PORT string
)

func main() {
	setup()

	app := app.Setup()
	log.Fatal(app.Start(fmt.Sprintf(":%s", PORT)))
}

func setup() {
	PORT = os.Getenv("PORT")
}
