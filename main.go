package main

import (
	"database/sql"
	"fmt"
	"learning_go/address"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "betainsitu"
	DB_PASSWORD = "betainsitu"
	DB_NAME     = "insitu_20170613"
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

	// fmt.Println(collection.Length())
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
	// getAllAddresses()

	br := testing.Benchmark(BenchmarkFunction)
	fmt.Println(br)
}
