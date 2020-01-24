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
	Servers     []models.Server
	DataSslLabs map[string]interface{}
}

// CallScannDomain Escanea el dominio pasado con ssllibas
func CallScannDomain(domain models.Domain) {
	scannDomainTruora := ScanerTruora{Domain: domain}
	scannDomainTruora.ScannDomain()
}

// SaveScannedDomain guardar
func (st *ScanerTruora) saveScannedDomain() {
	st.Domain.State = "R"
	models.SaveOrUpdateDomain((*st).Domain, true)
}

// ScannDomain Escanea el dominio pasado con ssllibas
func (st *ScanerTruora) ScannDomain() {
	bandera := 1
	ready := false
	var error error
	var status interface{}

	for bandera <= 150 && ready == false {
		fmt.Println("##### CallSslLabs #####")
		ready, (*st).DataSslLabs, status, error = CallSslLabs(st.Domain.DomainName)

		fmt.Println("Status: ", status)

		if error != nil {
			log.Println(error)
		}

		bandera = bandera + 1

		time.Sleep(time.Second * 10)
	}

	if ready == true {
		fmt.Println("Data: ", (*st).DataSslLabs)
		(*st).saveScannedDomain()
		(*st).processServers()
		// preprocessServers(data, domain)
	}
}

func (st *ScanerTruora) processServers() {
	fmt.Println("###Preprocesing###")
	servers := (*st).DataSslLabs["endpoints"].([]interface{})
	for i, u := range servers {
		fmt.Println(i, u)
		serverJSON := u.(map[string]interface{})
		server := models.NewServer(serverJSON["ipAddress"].(string))

		server.DomainID = (*st).Domain.ID
		err := models.FindStructBy(&server, "servers", "address", server.Address)

		if err != nil { //no lo encontro
			//Crearlo //si estado dominio esta en I primer escaneo, si esta en P lleva varios escaneos (servers_changed true)
			fmt.Println(server.Address)
			(*st).Servers = append((*st).Servers, server)
			models.SaveOrUpdateServer(server)
		} else {
			//Si lo encontro, imprimirlo
			(*st).Servers = append((*st).Servers, server)
			fmt.Println(server)
		}

	}
	fmt.Println("###EndPreprocesing###")
}

// if status == "IN_PROGRESS" {
// 	//fmt.Println("Guardar todos los servers")
// 	//ciclo y guardar
// }

// func preprocessServers(data map[string]interface{}, domain models.Domain) {
// 	fmt.Println("###Preprocesing###")
// 	servers := data["endpoints"].([]interface{})
// 	for i, u := range servers {
// 		fmt.Println(i, u)
// 		server := u.(map[string]interface{})
// 		s := models.NewServer(server["ipAddress"].(string))
// 		s.DomainID = domain.ID
// 		err := models.FindStructBy(&s, "servers", "address", s.Address)
// 		if err != nil { //no lo encontro
// 			//Crearlo //si estado dominio esta en I primer escaneo, si esta en P lleva varios escaneos (servers_changed true)
// 			fmt.Println(s.Address)
// 			models.SaveOrUpdateServer(s)
// 		} else {
// 			//Si lo encontro, imprimirlo
// 			fmt.Println(s)
// 		}
//
// 	}
// 	fmt.Println("###EndPreprocesing###")
// }

// fmt.Println("---------")
// fmt.Println(m)
// fmt.Println("---------")
// fmt.Println("#########")
// fmt.Println(m["status"])
// fmt.Println("#########")

// for k, v := range m {
//     switch vv := v.(type) {
//     case string:
//         fmt.Println(k, "is string", vv)
//     case float64:
//         fmt.Println(k, "is float64", vv)
//     case []interface{}:
//         fmt.Println(k, "is an array:")
//         for i, u := range vv {
//             fmt.Println(i, u)
//         }
//     default:
//         fmt.Println(k, "is of a type I don't know how to handle")
//     }
// }
