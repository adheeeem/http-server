package main

import (
	"net"
	"net/http"
	"os"
	"server/cmd/app"
	"server/pkg/banners"
)

func main() {
	host := "localhost"
	port := "9999"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(host string, port string) (err error) {
	mux := http.NewServeMux()
	bannersSvc := banners.NewService()
	server := app.NewServer(mux, bannersSvc)
	server.Init()
	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server,
	}
	return srv.ListenAndServe()
}
