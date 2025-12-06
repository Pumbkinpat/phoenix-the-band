package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/phoenix-the-band/internal/handler"
	"example.com/phoenix-the-band/internal/templates/layouts"
	"github.com/a-h/templ"
)

func main() {
	bandData, err := handler.LoadBandData("data/band.json")

	if err != nil {
		log.Fatal("Failed to load band data:", err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/static/css"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./web/public"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./web/public/img"))))

	component := layouts.BaseLayout(bandData)
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
