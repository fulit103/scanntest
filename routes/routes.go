package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fulit103/truoratest/models"
	"github.com/fulit103/truoratest/ssllabs"
	"github.com/go-chi/chi"
)

func InfoDomainEndPoint(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	fmt.Println("Domain: ", domain)
	go ssllabs.ScannDomain(domain)
	json.NewEncoder(w).Encode(&models.Domain{ServersChanged: "false", SslGrade: "A+", Title: domain})
}
