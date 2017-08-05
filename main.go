package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codingconcepts/env"
	"github.com/elazarl/goproxy"
)

func main() {
	config := struct {
		Port    int  `env:"PORT" required:"true"`
		Verbose bool `env:"VERBOSE" default:"false"`
	}{}

	if err := env.Set(&config); err != nil {
		log.Fatal(err)
	}

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = config.Verbose

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), proxy))
}
