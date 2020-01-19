package models

type Domain struct {
	ServersChanged   string   `json:"servers_changed,omitempty"`
	SslGrade         string   `json:"ssl_grade,omitempty"`
	PreviousSslGrade string   `json:"previous_ssl_grade,omitempty"`
	Logo             string   `json:"logo,omitempty"`
	Title            string   `json:"title,omitempty"`
	IsDown           string   `json:"is_down,omitempty"`
	Servers          []Server `json:"servers,omitempty"`
}

type Server struct {
	Address  string `json:"address,omitempty"`
	SslGrade string `json:"ssl_grade,omitempty"`
	Country  string `json:"country,omitempty"`
	Owner    string `json:"owner,omitempty"`
}

func saveDomain() {

}

func findDomain() {

}
