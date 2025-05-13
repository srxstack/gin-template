package main

import (
	"os"

	"github.com/srxstack/gin-template/cmd/gintpl-apiserver/app"

	_ "go.uber.org/automaxprocs"
)

func main() {
	command := app.NewGinTemplateCommand()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
