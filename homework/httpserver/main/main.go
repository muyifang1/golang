package main

import (
	"github.com/muyifang1/golang/homework/httpserver"
	"net/http"
)

func main() {
	//server := NewSdkHttpServer("test-server", MetricsFilterBuilder)
	//server := NewSdkHttpServer("server", LogFilterBuilder, TransferHeader)
	server := httpserver.NewSdkHttpServer("server", httpserver.LogFilterBuilder)

	server.Route(http.MethodPost, "/user/signup", httpserver.SignUp)

	server.Route(http.MethodGet, "/healthz", httpserver.Healthz)

	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
