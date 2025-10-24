package server

import (
	"net/http"
	"time"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
	idleTimeout  = 120 * time.Second
)

func New(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
}
