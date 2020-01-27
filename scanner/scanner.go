package scanner

import (
	"fmt"
	"log"
	"time"

	"github.com/fulit103/truoratest/models"
	whois "github.com/likexian/whois-go"
)

// ScanerTruora Clase para scanear dominio
type ScanerTruora struct {
	Domain      models.Domain
	DataSslLabs map[string]interface{}
}

// CallScannDomain Escanea el dominio pasado con ssllibas
func CallScannDomain(domain models.Domain) {
	scannDomainTruora := ScanerTruora{Domain: domain}
	scannDomainTruora.ScannDomain()
}

// ScannDomain Escanea el dominio pasado con ssllibas
func (st *ScanerTruora) ScannDomain() {
	bandera := 1
	ready := false
	var error error
	var status interface{}
	serversSaved := false

	go (*st).getLogo((*st).Domain)
	go (*st).getTitle((*st).Domain)

	for bandera <= 150 && ready == false {
		fmt.Println("_______________________")
		fmt.Println("_______________________")
		fmt.Println("##### CallSslLabs #####")
		ready, (*st).DataSslLabs, status, error = CallSslLabs((*st).Domain.DomainName)

		fmt.Println("Status: ", status)
		if status == "IN_PROGRESS" && serversSaved == false {
			endpoints := (*st).getEndpoints()
			(*st).saveServers(endpoints)
			(*st).getServersWhois(endpoints)
			serversSaved = true
		}

		if error != nil {
			log.Println(error)
		}

		bandera = bandera + 1

		time.Sleep(time.Second * 10)
		fmt.Println()
		fmt.Println()
	}

	(*st).Domain.State = "R"
	(*st).Domain.Updated = time.Now()
	domainDB := models.DomainDB{}

	if ready == true {
		//fmt.Println("Data: ", (*st).DataSslLabs)

		endpoints := (*st).getEndpoints()
		sslGrade := (*st).getServerGrade(endpoints)
		(*st).Domain.SslGrade = sslGrade
		(*st).saveServers(endpoints)

		domainDB.Update((*st).Domain, []string{"state", "ssl_grade", "updated"})
	} else {
		domainDB.Update((*st).Domain, []string{"state", "updated"})
	}
}

func (st *ScanerTruora) getEndpoints() []models.Server {
	endpoints := []models.Server{}
	servers := (*st).DataSslLabs["endpoints"].([]interface{})
	for i := range servers {
		//fmt.Println("Server endpoint: ", i, u)
		u := servers[i]
		serverJSON := u.(map[string]interface{})
		ipAddress := serverJSON["ipAddress"].(string)
		server := models.NewServer(ipAddress)
		if serverJSON["grade"] != nil {
			server.SslGrade = serverJSON["grade"].(string)
		}
		server.DomainID = (*st).Domain.ID
		endpoints = append(endpoints, server)
		fmt.Println("Server Add: ", server)
	}
	return endpoints
}

func (st *ScanerTruora) saveServers(servers []models.Server) {
	serverDB := models.ServerDB{}
	for i := range servers {
		s, err := serverDB.FindBy(servers[i].Address, servers[i].DomainID)
		if err != nil {
			serverDB.SaveOrUpdate(servers[i])
		} else {
			if servers[i].SslGrade != "" {
				s.SslGrade = servers[i].SslGrade
			}
			serverDB.SaveOrUpdate(s, true)
			fmt.Println("Server encontrado: ", s)
		}
	}
}

func (st *ScanerTruora) getServerGrade(servers []models.Server) string {
	for i := range servers {
		return servers[i].SslGrade
	}
	return ""
}

func (st *ScanerTruora) getTitle(domain models.Domain) {
	title, err := GetHTMLTitleFromURL("http://" + domain.DomainName)
	if err == nil {
		domainDB := models.DomainDB{}
		domain.Title = title
		domainDB.Update(domain, []string{"title"})
		fmt.Println("Title: " + title)
	}
}

func (st *ScanerTruora) getLogo(domain models.Domain) {
	logo, err := GetHTMLLogoFromURL("http://" + domain.DomainName)
	if err == nil {
		fmt.Println("Logo: " + logo)
		domainDB := models.DomainDB{}
		domain.Logo = logo
		domainDB.Update(domain, []string{"logo"})
	}
}

func (st *ScanerTruora) getServersWhois(servers []models.Server) {
	for _, s := range servers {
		go (*st).getWhois(s)
	}
}

func (st *ScanerTruora) getWhois(server models.Server) {
	result, err := whois.Whois(server.Address)
	if err == nil {
		fmt.Println("-----")
		name, _ := getWhoisField(result, "OrgName:")
		fmt.Println("OrgName: ", name)
		contry, _ := getWhoisField(result, "Country:")
		fmt.Println("Contry: ", contry)
		fmt.Println("-----")
	}
}
