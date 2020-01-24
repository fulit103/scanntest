package models

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

func connect() (*sqlx.DB, error) {
	//addr := os.Getenv("DB")
	//fmt.Println("Postgres addr: " + addr)
	db, err := sqlx.Connect("postgres", "postgresql://root@0.0.0.0:26257/truora?sslmode=disable")
	if err != nil {
		fmt.Println("Could not connect...")
	} else {
		fmt.Println("Connecting successful")
	}
	return db, err
}

// PrepareValues get array convert in format INSERT ':val1,:val2,:val3'
func PrepareValues(array []string, tipo string) string {
	values := []string{}
	for index := range array {
		if tipo == "INSERT" {
			values = append(values, ":"+array[index])
		} else {
			values = append(values, array[index]+"=:"+array[index])
		}
	}
	return strings.Join(values, ",")
}

func saveOrUpdateStruct(object interface{}, table string, fields []string, pk string, forceUpdateOptional ...bool) {
	forceUpdate := false
	if len(forceUpdateOptional) > 0 {
		forceUpdate = forceUpdateOptional[0]
	}

	db, err := connect()
	if err == nil {

		if forceUpdate == false {
			query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, table, strings.Join(fields, ","), PrepareValues(fields, "INSERT"))
			fmt.Println(query)

			_, err := db.NamedExec(query, object)
			if err != nil {
				log.Println("Error: ", err)
			}
		} else {
			query := fmt.Sprintf(`UPDATE %s SET %s WHERE %s=%s`, table, PrepareValues(fields, "UPDATE"), pk, ":"+pk)
			fmt.Println(query)

			_, err := db.NamedExec(query, object)
			if err != nil {
				log.Println("Error: ", err)
			}
		}
	}
}

// FindStructBy find
func FindStructBy(object interface{}, table string, field string, value string) error {
	db, err := connect()

	if err == nil {
		sql := fmt.Sprintf(`SELECT * FROM %s WHERE %s='%s' `, table, field, value)
		fmt.Println("SQL: ", sql)
		rows, _ := db.Queryx(sql)
		for rows.Next() {
			errStruct := rows.StructScan(object)
			if errStruct != nil {
				fmt.Println("error aqui: ", errStruct)
			}
			return nil
		}
	}
	return errors.New("not found")
}

// FindAllStruct listar todos los elementos de una tabla
func FindAllStruct(dest interface{}, table string, page int, limit int) error {
	db, err := connect()

	if err == nil {
		arr := reflect.ValueOf(dest).Elem()
		v := reflect.New(reflect.TypeOf(dest).Elem().Elem())

		sql := fmt.Sprintf(`SELECT * FROM %s LIMIT %d OFFSET %d;`, table, limit, page)

		rows, err := db.Queryx(sql)
		if err == nil {
			for rows.Next() {
				if err = rows.StructScan(v.Interface()); err == nil {
					arr.Set(reflect.Append(arr, v.Elem()))
				} else {
					log.Println(err)
				}
			}
		} else {
			log.Println(err)
		}
		return nil
	}
	return err
}