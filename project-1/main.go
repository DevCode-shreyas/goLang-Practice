package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var products []Product

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: homePage")
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnAllProducts")
	json.NewEncoder(w).Encode(products)
}

func getproduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	// log.Println(key)
	for _, product := range products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", returnAllProducts)
	myRouter.HandleFunc("/product/{id}", getproduct)
	myRouter.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:10000", nil)
}

func cheackError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type Data struct {
	Id   int
	Name string
}

func main() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBUser, DBPassword, DBName)

	db, err := sql.Open("mysql", connectionString)
	cheackError(err)
	defer db.Close()

	result, err := db.Exec("insert into data values(4,'xyz')")
	cheackError(err)
	lastInsertedId, err := result.LastInsertId()
	fmt.Println("lastInsertedId", lastInsertedId)
	cheackError(err)
	rowsAffected, err := result.RowsAffected()
	fmt.Println("rowsAffected", rowsAffected)
	cheackError(err)

	rows, err := db.Query("SELECT * FROM data")
	cheackError(err)

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.Id, &data.Name)
		cheackError(err)
		fmt.Println(data)
	}

	products = []Product{
		{Id: "1", Name: "Chair", Quantity: 100, Price: 100.00},
		{Id: "2", Name: "Desk", Quantity: 200, Price: 200.00},
	}
	handleRequests()
}
