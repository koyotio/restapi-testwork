package main

import (
	"github.com/valyala/fasthttp"
	"time"
)

type Server struct {
	httpServer *fasthttp.Server
}

func (s *Server) Run(listenAddr string, handler fasthttp.RequestHandler) error {
	s.httpServer = &fasthttp.Server{
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s.httpServer.ListenAndServe(listenAddr)
}
