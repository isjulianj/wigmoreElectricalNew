package main

import (
	"fmt"
	"net/http"
)

func MuxHandler(app *Application) http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir(fmt.Sprintf("./%s", app.Config.AssetsDir)))
	mux.Handle("GET /static/*", http.StripPrefix(fmt.Sprintf("/%s", app.Config.AssetsDir), fileServer))

	// home
	mux.HandleFunc("GET /{$}", HomePage())

	////about
	//mux.HandleFunc("GET /about", handlers.AboutHandler(app))
	//
	//// blog related
	//mux.HandleFunc("GET /blog/{$}", handlers.BlogHandler(app))
	//mux.HandleFunc("GET /blog/post/{slug}", handlers.PostPageHandler(app))
	//mux.HandleFunc("/create/post/{$}", handlers.CreatePostHandler(app))
	//
	//// author related
	//mux.HandleFunc("GET /author/{slug}", handlers.AuthorHandler(app))
	//
	////helpers
	//mux.HandleFunc("GET /images/user/{imageId}", handlers.GetUserImage(app))

	return mux

}
