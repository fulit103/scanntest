package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fulit103/truoratest/models"
	"github.com/fulit103/truoratest/scanner"
	"github.com/go-chi/chi"
)

// InfoDomainEndPoint url paso 1
func InfoDomainEndPoint(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")
	fmt.Println("Domain: ", domain)
	//models.SaveDomain(domain)
	//d := models.FindDomainBy("domain", domain)
	d := models.Domain{}

	err := models.FindStructBy(&d, "domains", "domain", domain)
	if err == nil {
		fmt.Println("Domain Find: ", d)
		if d.State == "R" {
			go scanner.CallScannDomain(d)
		} else {
			fmt.Println("--------")
			fmt.Println("Ya esta ejecutandose el proceso")
			fmt.Println("--------")
		}
	} else {
		d := models.NewDomain(domain)
		models.SaveOrUpdateDomain(d)
		models.FindStructBy(&d, "domains", "domain", domain)
		go scanner.CallScannDomain(d)
	}

	json.NewEncoder(w).Encode(&models.Domain{ServersChanged: "false", SslGrade: "A+", Title: domain, DomainName: domain})
}

// ListDomainsEndPoint listar dominios consultados: paso 2
func ListDomainsEndPoint(w http.ResponseWriter, r *http.Request) {
	//listado := map[string] []interface{"servidores": 11, "dominios": 22}
	listado := make(map[string]interface{})

	data := []models.Domain{}
	models.FindAllStruct(&data, "domains", 0, 10)

	listado["dominios"] = data

	fmt.Println(listado)

	//&models.Domain{ServersChanged: "false", SslGrade: "A+", Title: domain, DomainName: domain}

	json.NewEncoder(w).Encode(&listado)
}
