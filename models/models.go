package models

import (
	"time"

	_ "github.com/lib/pq"
)

// Domain encapsula instancia principal
type Domain struct {
	ID               string    `json:"id,omitempty" db:"id"`
	ServersChanged   string    `json:"servers_changed,omitempty" db:"servers_changed"`
	SslGrade         string    `json:"ssl_grade" db:"ssl_grade"`
	PreviousSslGrade string    `json:"previous_ssl_grade,omitempty" db:"previous_ssl_grade"`
	Logo             string    `json:"logo,omitempty" db:"logo"`
	Title            string    `json:"title,omitempty" db:"title"`
	IsDown           bool      `json:"is_down,omitempty" db:"is_down"`
	DomainName       string    `json:"domain" db:"domain"`
	LastCall         string    `json:"last_call" db:"last_call"`
	Created          string    `json:"created" db:"created"`
	Updated          time.Time `json:"updated" db:"updated"`
	State            string    `json:"state" db:"state"` // I : initial processing, P processing, R ready
}

// NewDomain crea una instancia de Domain.
func NewDomain(domain string) Domain {
	return Domain{
		DomainName:       domain,
		ServersChanged:   "",
		SslGrade:         "",
		PreviousSslGrade: "",
		Logo:             "",
		Title:            "",
		IsDown:           false,
		State:            "I"}
}

// Server es la estructura para almacenar la info de los servidores
type Server struct {
	ID       string `json:"id,omitempty" db:"id" db:"id"`
	Address  string `json:"address,omitempty" db:"address"`
	SslGrade string `json:"ssl_grade,omitempty" db:"ssl_grade"`
	Country  string `json:"country,omitempty" db:"country"`
	Owner    string `json:"owner,omitempty" db:"owner"`
	InUse    bool   `json:"in_use,omitempty"`
	DomainID string `json:"domain,omitempty" db:"domain_id"`
	Created  string `json:"created" db:"created"`
	Updated  string `json:"updated" db:"updated"`
}

// NewServer crea una instancia de tipo Server
func NewServer(address string) Server {
	return Server{
		Address:  address,
		SslGrade: "",
		Country:  "",
		Owner:    "",
		InUse:    true}
}
