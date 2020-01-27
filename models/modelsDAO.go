package models

import "fmt"

// DomainDB objeto para el almacenamiento y busqueda en la bd
type DomainDB struct {
}

// SaveOrUpdate guarda o actualiza el dominio en la base de datos
func (dB *DomainDB) SaveOrUpdate(d Domain, forceUpdateOptional ...bool) {
	forceUpdate := false
	if len(forceUpdateOptional) > 0 {
		forceUpdate = forceUpdateOptional[0]
	}
	fields := []string{"domain", "ssl_grade", "previous_ssl_grade", "logo", "title", "is_down", "state"}
	//models.SaveOrUpdateDomain(domain, false)
	saveOrUpdateStruct(&d, "domains", fields, "domain", forceUpdate)
}

// Update actualizar datos de un dominio
func (dB *DomainDB) Update(d Domain, columnas []string) {
	saveOrUpdateStruct(&d, "domains", columnas, "domain", true)
}

// FindBy busa un dominio en la base de datos, por el dominio
func (dB *DomainDB) FindBy(domain string) (Domain, error) {
	d := NewDomain("")
	err := FindStructBy(&d, "domains", fmt.Sprintf("domain='%s'", domain))
	if err == nil {
		return d, nil
	}
	return d, err
}

// ServerDB encapsula los metodos para el almacenamiento y busqueda de la tabla Server
type ServerDB struct {
}

// SaveOrUpdate crear o actualizar tabla server
func (sS *ServerDB) SaveOrUpdate(s Server, forceUpdateOptional ...bool) {
	forceUpdate := false
	if len(forceUpdateOptional) > 0 {
		forceUpdate = forceUpdateOptional[0]
	}
	fields := []string{"address", "ssl_grade", "country", "owner", "domain_id"}
	//models.SaveOrUpdateDomain(domain, false)
	saveOrUpdateStruct(&s, "servers", fields, "id", forceUpdate)
}

// FindBy busa un dominio en la base de datos, por el dominio
func (sS *ServerDB) FindBy(ip string, domainID string) (Server, error) {
	s := NewServer("")
	err := FindStructBy(&s, "servers", fmt.Sprintf("address='%s' AND domain_id='%s'", ip, domainID))
	if err == nil {
		return s, nil
	}
	return s, err
}

// FindAllBy busca los servidores de un dominio
func (sS *ServerDB) FindAllBy(domainID string) []Server {
	servers := []Server{}
	FindAllStruct(&servers, "servers", 0, 10, fmt.Sprintf("domain_id=%s", domainID))
	return servers
}

// Update actualizar datos de un server
func (sS *ServerDB) Update(s Server, columnas []string) {
	saveOrUpdateStruct(&s, "servers", columnas, "id", true)
}
