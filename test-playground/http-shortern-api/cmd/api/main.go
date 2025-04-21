package main

import (
	"errors"
	"http-shortern-api/internal/handlers"
	"http-shortern-api/internal/store"
	"log/slog"
	"net/http"
	"time"
)

const (
	addr    = "localhost:8080"
	timeout = 10 * time.Second
)

func main() {
	log := slog.With("app", "urlshortern")

	slog.SetDefault(log)

	log.Info("Starting", "addr", addr)

	// urls := func(w http.ResponseWriter, r *http.Request) {
	// 	_, _ = fmt.Fprintln(w, "hello from the bite links server!")
	// }

	server := handlers.NewServer(store.NewStore())
	timeoutServer := http.TimeoutHandler(server, timeout, "timeout")

	srv := &http.Server{
		Addr:         addr,
		Handler:      timeoutServer,
		ReadTimeout:  timeout * 2,
		WriteTimeout: timeout * 4,
	}
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		log.Error("Server closed unexpectedly", "message", err)
	}
}
