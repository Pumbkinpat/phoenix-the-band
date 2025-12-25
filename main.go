package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/phoenix-the-band/internal/handler"
	"github.com/phoenix-the-band/internal/templates/pages"

	"github.com/a-h/templ"
)

func main() {
	bandData, err := handler.LoadBandData("data/band.json")

	if err != nil {
		log.Fatal("Failed to load band data:", err)
	}

	component := pages.HomePage(bandData)

	mux := http.NewServeMux()

	mux.Handle("/", http.RedirectHandler("/home", http.StatusTemporaryRedirect))
	mux.Handle("/home", templ.Handler(component))

	mux.Handle("/public/", http.StripPrefix("/public/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		http.FileServer(http.Dir("./web/public")).ServeHTTP(w, r)
	})))

	mux.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		http.FileServer(http.Dir("./web/static")).ServeHTTP(w, r)
	})))

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	fmt.Println("Listening on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

// make sure to use the correct path
// to the folder/file we want to server
// or it will causes MIME error in dev tool

// changed file might not be updated to browser
// because of caching feature, we can avoid this by
// hiting hard reset shortcut in chrome `ctr + shift + r`
// of disable caching in network tab in dev tool
