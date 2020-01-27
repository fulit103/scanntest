package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fulit103/truoratest/models"
	"github.com/fulit103/truoratest/scanner"
	"github.com/go-chi/chi"
)

// InfoDomainEndPoint url paso 1
func InfoDomainEndPoint(w http.ResponseWriter, r *http.Request) {
	domain := chi.URLParam(r, "domain")

	domainDB := models.DomainDB{}

	d, err := domainDB.FindBy(domain)
	if err == nil {
		if d.State == "R" {
			d.State = "P"
			d.Updated = time.Now()
			domainDB.Update(d, []string{"state", "updated"})
			go scanner.CallScannDomain(d)
		} else {
			fmt.Println("Ya esta ejecutandose el proceso")
		}
	} else {
		d = models.NewDomain(domain)
		domainDB.SaveOrUpdate(d)
		d, _ = domainDB.FindBy(domain)
		go scanner.CallScannDomain(d)
	}

	serverDB := models.ServerDB{}
	servers := serverDB.FindAllBy(d.ID)
	serializer := models.DomainSerializer{Domain: d, Servers: servers}

	json.NewEncoder(w).Encode(serializer.GetData())
}

// ListDomainsEndPoint listar dominios consultados: paso 2
func ListDomainsEndPoint(w http.ResponseWriter, r *http.Request) {
	domains := []models.Domain{}
	models.FindAllStruct(&domains, "domains", 0, 10, "true", "updated DESC")

	servers := []models.Server{}
	models.FindAllStruct(&servers, "servers", 0, 300, models.GetIds(domains))

	serializer := models.DomainSerializer{Domains: domains, Servers: servers}

	json.NewEncoder(w).Encode(serializer.GetDataMany())
}
