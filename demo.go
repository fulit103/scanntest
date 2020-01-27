package main

// func main() {
// 	resp, err := http.Get("https://www.rapigo.co")
//
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
//
// 	if title, ok, _ := scanner.GetHTMLTitle(resp.Body); ok {
// 		println(title)
// 	} else {
// 		println("Fail to get HTML title")
// 	}
// }

// func main() {
// 	m := models.FindStructBy(models.Domain{}, "domains", []string{"domain"}, []interface{"www.rapigo.co", 1})
// 	fmt.Println("Domain: ", m)
// }

// func main() {
// 	// models := models.FindStrucAll(models.Domain{}, "domains", 1, 10).([10]models.Domain)
// 	// fmt.Println(models)
// 	data := []models.Domain{}
// 	models.FindAllStruct(&data, "domains", 0, 10)
// 	fmt.Println(data)
// }

// import (
// 	"fmt"
// 	"reflect"
// )
//
// func createElement(t reflect.Type, args ...interface{}) interface{} {
// 	//a := reflect.New(t)
// 	//a := reflect.Zero(reflect.SliceOf(t)).Interface()
// 	a := reflect.ArrayOf(len(args), t)
// 	array := reflect.New(a).Elem()
//
// 	fmt.Println("R: ", a)
// 	if len(args) > 0 {
// 		for i := 0; i < len(args); i++ {
// 			//a = append(a, reflect.ValueOf(args[i]))
// 			//a[i] = reflect.ValueOf(args[i])
// 			array.Index(i).Set(reflect.ValueOf(args[i]))
// 		}
// 	}
//
// 	return array.Interface()
// }
//
// func main() {
// 	var i float64
//
// 	a := createElement(reflect.TypeOf(i), 1.2, 2.3, 3.5).([3]float64)
//
// 	fmt.Println(a[0], a[1], a[2])
//
// }
