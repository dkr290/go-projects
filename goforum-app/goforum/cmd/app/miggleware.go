package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/justinas/nosurf"
)

func LogRequestInfo(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Printf("%d/%d/%d : %d:%d\n", now.Month(), now.Day(), now.Year(), now.Hour(), now.Minute())
		fmt.Println("Url path :", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func SetupSession(next http.Handler) http.Handler {
	return sm.LoadAndSave(next)
}

func NoSurf(next http.Handler) http.Handler {
	nosurfHandler := nosurf.New(next)
	nosurfHandler.SetBaseCookie(http.Cookie{
		Name:     "mycsrfcookie",
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
		MaxAge:   3600,
		SameSite: http.SameSiteLaxMode,
	})

	return nosurfHandler
}