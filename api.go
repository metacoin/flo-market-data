package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type APIgetAllResponse struct {
	Unixtime     int
	Cryptsy      string
	Poloniex     string
	Bittrex      string
	Daily_volume string
	USD          string
}

func APIgetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	method := r.Method
	if method == "" {
		method = "GET"
	}
	fmt.Println(method + " " + r.URL.Path + " " + r.RemoteAddr)

	dbtx, err := DBH.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := dbtx.Prepare(`select unixtime, cryptsy, poloniex, bittrex, daily_volume, USD from markets order by unixtime desc LIMIT 1`)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	var resultCount int

	m := APIgetAllResponse{}
	for rows.Next() {
		rows.Scan(&m.Unixtime, &m.Cryptsy, &m.Poloniex, &m.Bittrex, &m.Daily_volume, &m.USD)
		resultCount++
		break
	}
	if resultCount < 1 {
		fmt.Printf("no results from query...")
	}

	dbtx.Commit()
	stmt.Close()
	rows.Close()

	json, err := json.Marshal(m)
	if err != nil {
		fmt.Println("exit code 300")
		log.Fatal(err)
	}

	w.Write(json)

}
