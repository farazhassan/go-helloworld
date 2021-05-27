package main

import (
	"fmt"
	"net/http"

	handlers "takehome/handlers"
	middlewares "takehome/middlewares"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port uint16 `envconfig:"PORT" required:"true"`
}

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		panic(err.Error())
	}

	router := http.NewServeMux()

	router.Handle("/echo", handlers.RootHandler(handlers.Echo))
	router.Handle("/invert", handlers.RootHandler(handlers.Invert))
	router.Handle("/multiply", handlers.RootHandler(handlers.Multiply))
	router.Handle("/flatten", handlers.RootHandler(handlers.Flatten))
	router.Handle("/sum", handlers.RootHandler(handlers.Sum))

	http.ListenAndServe(fmt.Sprintf(":%d", c.Port), middlewares.NewFileToMatrixMiddleware(middlewares.NewPOSTMethodOnlyMiddleware(router)))
}
