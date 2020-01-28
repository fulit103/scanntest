package models

var schema = `
CREATE TABLE IF NOT EXISTS domains (
  id SERIAL PRIMARY KEY,
  domain varchar(30) NOT NULL,
  ssl_grade varchar(3) NOT NULL,
  previous_ssl_grade varchar(3) NOT NULL,
  logo text NOT NULL,
  title text NOT NULL,
  is_down boolean default false,
  state char(1) NOT NULL,
  last_call timestamp DEFAULT NOW(),
  created timestamp DEFAULT NOW(),
  updated timestamp DEFAULT NOW()

);

CREATE TABLE  IF NOT EXISTS servers (
  id SERIAL PRIMARY KEY,
  address varchar(30) NOT NULL,
  ssl_grade varchar(3) NOT NULL,
  country varchar(2) NOT NULL,
  owner varchar(200) NOT NULL,
  in_use boolean default true,
  created timestamp DEFAULT NOW(),
  updated timestamp DEFAULT NOW(),
  domain_id INT REFERENCES domains(id)
)`

// InitDB Crea las tablas en el sistema
func InitDB() {
	db, err := connect()
	if err == nil {
		db.MustExec(schema)
		db.Exec("UPDATE domains SET state='R' where true")
	}
}
