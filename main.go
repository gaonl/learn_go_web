package main

import (
	"config"
	"fmt"
	"net/http"
	"route"
	"time"
)

func createMuxAndRegisterHandlers() (*http.ServeMux) {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.AppConfig.ServerConfig.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", route.Index)
	mux.HandleFunc("/err", route.Err)

	mux.HandleFunc("/login", route.Login)
	mux.HandleFunc("/logout", route.Logout)
	mux.HandleFunc("/signup", route.Signup)
	mux.HandleFunc("/signup_account", route.SignupAccount)
	mux.HandleFunc("/authenticate", route.Authenticate)

	mux.HandleFunc("/thread/new", route.NewThread)
	mux.HandleFunc("/thread/create", route.CreateThread)
	mux.HandleFunc("/thread/post", route.PostThread)
	mux.HandleFunc("/thread/read", route.ReadThread)

	return mux
}

func main() {
	fmt.Println("Chitchat", config.AppConfig.ServerConfig.Version, "started at", config.AppConfig.ServerConfig.Address)
	config.Info("Server Started")

	mux := createMuxAndRegisterHandlers()

	server := &http.Server{
		Addr:           config.AppConfig.ServerConfig.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.AppConfig.ServerConfig.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.AppConfig.ServerConfig.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
