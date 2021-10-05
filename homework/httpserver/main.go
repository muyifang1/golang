package main

import "net/http"

func main() {
	//server := NewSdkHttpServer("test-server", MetricsFilterBuilder)
	//server := NewSdkHttpServer("server", LogFilterBuilder, TransferHeader)
	server := NewSdkHttpServer("server", LogFilterBuilder)

	server.Route(http.MethodPost, "/user/signup", SignUp)

	server.Route(http.MethodGet, "/healthz", healthz)

	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
