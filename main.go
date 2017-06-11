package main

import (
	"fmt"
	"./address"
	"database/sql"
	_ "github.com/lib/pq"
	"testing"
)



const (
	DB_USER     = "user"
	DB_PASSWORD = "password"
	DB_NAME     = "database"
)

func getAllAddresses() {


	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	//fmt.Println("# Querying")
	rows, err := db.Query("SELECT id, recipient_additional, street, street_additional, postcode, city, company_additional, lat, lng, company_label, region FROM addresses")
	checkErr(err)

	var collection address.AddressCollection


	for rows.Next() {
		var instance = address.AddressInstance{}
		rows.Scan(
			&instance.Id,
			&instance.Recipient_additional,
			&instance.Street,
			&instance.Street_additional,
			&instance.Postcode,
			&instance.City,
			&instance.Company_additional,
			&instance.Lat,
			&instance.Lng,
			&instance.Company_label,
			&instance.Region,
		)
		collection.Add(instance)
	}

	fmt.Println(collection.Get(850))

	fmt.Println(collection.Length())
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func BenchmarkFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {

		getAllAddresses()
	}
}

func main() {
	//getAllAddresses()

	var mockedInstance address.AddressInstance

	mockedInstance.Id = 12
	mockedInstance.Recipient_additional = ""
	mockedInstance.Street = "Teststree 12"
	mockedInstance.Street_additional = ""
	mockedInstance.Postcode = "12345"
	mockedInstance.City = "Berlin"
	mockedInstance.Company_additional = ""
	mockedInstance.Lat = 0
	mockedInstance.Lng = 0
	mockedInstance.Company_label = ""
	mockedInstance.Region = ""


	var mockedCollection address.AddressCollection
	mockedCollection.Add(mockedInstance)

	fmt.Println(mockedCollection.Get(12))

	//br := testing.Benchmark(BenchmarkFunction)
	//fmt.Println(br)
}
