package scanner

import (
	"fmt"
	"log"
	"time"

	"github.com/fulit103/truoratest/models"
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

// SaveScannedDomain guardar
func (st *ScanerTruora) saveScannedDomain() {
	(*st).Domain.State = "R"
	domainDB := models.DomainDB{}
	domainDB.SaveOrUpdate((*st).Domain, true)
}

// ScannDomain Escanea el dominio pasado con ssllibas
func (st *ScanerTruora) ScannDomain() {
	bandera := 1
	ready := false
	var error error
	var status interface{}
	serversSaved := false

	for bandera <= 150 && ready == false {
		fmt.Println("_______________________")
		fmt.Println("_______________________")
		fmt.Println("##### CallSslLabs #####")
		ready, (*st).DataSslLabs, status, error = CallSslLabs(st.Domain.DomainName)

		fmt.Println("Status: ", status)
		if status == "IN_PROGRESS" && serversSaved == false {
			endpoints := (*st).getEndpoints()
			(*st).saveServers(endpoints)
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

	if ready == true {
		//fmt.Println("Data: ", (*st).DataSslLabs)

		endpoints := (*st).getEndpoints()
		sslGrade := (*st).getServerGrade(endpoints)
		(*st).Domain.SslGrade = sslGrade
		(*st).saveServers(endpoints)
		(*st).saveScannedDomain()
	} else {
		(*st).saveScannedDomain()
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
			s.SslGrade = servers[i].SslGrade
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
