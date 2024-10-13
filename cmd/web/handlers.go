package main

import (
	homeindex "github.com/isjulianj/mutuo-go/ui/home"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func HomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		homeindex.Show().Render(r.Context(), w)
	}
}
