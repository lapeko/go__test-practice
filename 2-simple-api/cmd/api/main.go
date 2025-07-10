package main

import (
	"github.com/lapeko/go__test-practice/2-simple-api/internal/api"
	"github.com/lapeko/go__test-practice/2-simple-api/internal/logger"
)

func main() {
	a := api.New()
	a.SetupMiddlewares()
	a.SetupRouter()
	logger.Log.Fatal(a.Start())
}
