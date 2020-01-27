package models

// DomainSerializer clase para serializar
type DomainSerializer struct {
	Domain  Domain
	Domains []Domain
	Servers []Server
}

// GetDataMany return Json
func (ds *DomainSerializer) GetDataMany() map[string]interface{} {
	array := make(map[string]interface{})
	listDomains := make([]interface{}, 0)
	if (*ds).Domains != nil && (*ds).Servers != nil {
		for i := 0; i < len((*ds).Domains); i++ {
			domain := (*ds).Domains[i]

			serversOfDomine := make([]Server, 0)
			for j := 0; j < len((*ds).Servers); j++ {
				if domain.ID == (*ds).Servers[j].DomainID {
					serversOfDomine = append(serversOfDomine, (*ds).Servers[j])
				}
			}

			object := (*ds).getJSONRow(domain, serversOfDomine)

			listDomains = append(listDomains, object)
		}
		array["items"] = listDomains
	}

	return array
}

// GetData return JSON row
func (ds *DomainSerializer) GetData() interface{} {
	object := (*ds).getJSONRow((*ds).Domain, (*ds).Servers)
	return object
}

func (ds *DomainSerializer) getJSONRow(domain Domain, servers []Server) map[string]interface{} {
	object := make(map[string]interface{})

	listServers := make([]interface{}, 0)
	for j := 0; j < len(servers); j++ {
		s := servers[j]
		serverJSON := make(map[string]interface{})
		serverJSON["address"] = s.Address
		serverJSON["ssl_grade"] = s.SslGrade
		serverJSON["country"] = s.Country
		serverJSON["owner"] = s.Owner
		listServers = append(listServers, serverJSON)
	}

	object["servers"] = listServers

	object["servers_changed"] = domain.ServersChanged
	object["ssl_grade"] = domain.SslGrade
	object["previous_ssl_grade"] = domain.PreviousSslGrade
	object["logo"] = domain.Logo
	object["title"] = domain.Title
	object["is_down"] = domain.IsDown
	object["domain"] = domain.DomainName

	return object
}
