package main

import (
	"fmt"
	"net/http"

	"github.com/fulit103/truoratest/routes"
	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Hola mundo")
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Truora api by Julian Toro"))
	})

	r.Get("/analyse/{domain}", routes.InfoDomainEndPoint)
	http.ListenAndServe(":3000", r)
}
