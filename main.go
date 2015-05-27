package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	DB_DIR  = "db"
	DB_PATH = "./" + DB_DIR + "/markets.db"
)

type API_URLs struct {
	Bitcoinaverage struct {
		USDBTCLAST string `json:"USD_BTC_LAST"`
	} `json:"bitcoinaverage"`
	Bittrex struct {
		BTCFLOLAST string `json:"BTC_FLO_LAST"`
		BTCFLOVOL  string `json:"BTC_FLO_VOL"`
	} `json:"bittrex"`
	Cryptsy struct {
		BTCLTCLAST string `json:"BTC_LTC_LAST"`
		LTCFLOLAST string `json:"LTC_FLO_LAST"`
		LTCFLOVOL  string `json:"LTC_FLO_VOL"`
	} `json:"cryptsy"`
	Poloniex struct {
		BTCFLOLAST string `json:"BTC_FLO_LAST"`
		BTCFLOVOL  string `json:"BTC_FLO_VOL"`
	} `json:"poloniex"`
}

type MarketData struct {
	bittrex_BTC_FLO_last  float64
	bittrex_BTC_FLO_volu  float64
	poloniex_BTC_FLO_last float64
	poloniex_BTC_FLO_volu float64
	cryptsy_BTC_LTC_last  float64
	cryptsy_LTC_FLO_last  float64
	cryptsy_LTC_FLO_volu  float64
	bitcoinaverage_USD    float64

	FLO_24h_vol                float64
	bittrex_BTC_FLO_vol_share  float64
	poloniex_BTC_FLO_vol_share float64
	cryptsy_LTC_FLO_vol_share  float64
	BTC_FLO_last_weighted      float64
	USD_FLO_last_weighted      float64
}

func initHTTP() {

	// init http handlers from api.go
	http.HandleFunc("/flo-market-data/v1/getAll", APIgetAll)

	// start listening on port 41290
	log.Println("Listening on port 41290")
	err := http.ListenAndServe(":41290", nil)
	if err != nil {
		log.Fatal("ListenAndServe failure: ", err)
		fmt.Printf("%v", err.Error())
	}

}

func main() {
	go watchMarkets()
	initHTTP()
}

func watchMarkets() {
	initDB()
	createTables()

	conf := getConfig()
	market_data := MarketData{}

	for {
		// bittrex
		market_data.bittrex_BTC_FLO_last = get_bittrex_btc_flo_last(conf.Bittrex.BTCFLOLAST)
		market_data.bittrex_BTC_FLO_volu = get_bittrex_btc_flo_volu(conf.Bittrex.BTCFLOLAST)

		// poloniex
		market_data.poloniex_BTC_FLO_last = get_poloniex_btc_flo_last(conf.Poloniex.BTCFLOLAST)
		market_data.poloniex_BTC_FLO_volu = get_poloniex_btc_flo_volu(conf.Poloniex.BTCFLOVOL)

		// cryptsy
		market_data.cryptsy_BTC_LTC_last = get_cryptsy_btc_ltc_last(conf.Cryptsy.BTCLTCLAST)
		market_data.cryptsy_LTC_FLO_last, market_data.cryptsy_LTC_FLO_volu = get_cryptsy_ltc_flo_last(conf.Cryptsy.LTCFLOLAST)

		// bitcoinaverage
		market_data.bitcoinaverage_USD = get_bitcoinaverage_usd(conf.Bitcoinaverage.USDBTCLAST)

		// calculate stuff
		market_data.FLO_24h_vol = market_data.bittrex_BTC_FLO_volu + market_data.poloniex_BTC_FLO_volu + market_data.cryptsy_LTC_FLO_volu
		market_data.bittrex_BTC_FLO_vol_share = market_data.bittrex_BTC_FLO_volu / market_data.FLO_24h_vol
		market_data.poloniex_BTC_FLO_vol_share = market_data.poloniex_BTC_FLO_volu / market_data.FLO_24h_vol
		market_data.cryptsy_LTC_FLO_vol_share = market_data.cryptsy_LTC_FLO_volu / market_data.FLO_24h_vol

		market_data.BTC_FLO_last_weighted = (market_data.bittrex_BTC_FLO_vol_share * market_data.bittrex_BTC_FLO_last) + (market_data.poloniex_BTC_FLO_vol_share * market_data.poloniex_BTC_FLO_last) + (market_data.cryptsy_LTC_FLO_vol_share * market_data.cryptsy_BTC_LTC_last * market_data.cryptsy_LTC_FLO_last)

		market_data.USD_FLO_last_weighted = market_data.bitcoinaverage_USD * market_data.BTC_FLO_last_weighted

		/*
			fmt.Printf("24hr volume: %7.8f\n\n", market_data.FLO_24h_vol)
			fmt.Printf("bittrex vol: %3.2f \n", 100*market_data.bittrex_BTC_FLO_vol_share)
			fmt.Printf("poloniex vl: %3.2f \n", 100*market_data.poloniex_BTC_FLO_vol_share)
			fmt.Printf("cryptsy vol: %3.2f \n\n", 100*market_data.cryptsy_LTC_FLO_vol_share)
			fmt.Printf("weighted   : %7.8f\n", market_data.BTC_FLO_last_weighted)
			fmt.Printf("flo/USD    : %7.8f\n", market_data.USD_FLO_last_weighted)
		*/

		bittrex_share := strconv.FormatFloat(market_data.bittrex_BTC_FLO_vol_share, 'f', 4, 32)
		poloniex_share := strconv.FormatFloat(market_data.poloniex_BTC_FLO_vol_share, 'f', 4, 32)
		cryptsy_share := strconv.FormatFloat(market_data.cryptsy_LTC_FLO_vol_share, 'f', 4, 32)
		daily_volume := strconv.FormatFloat(market_data.FLO_24h_vol, 'f', 8, 32)
		weighted := strconv.FormatFloat(market_data.BTC_FLO_last_weighted, 'f', 8, 32)
		USD := strconv.FormatFloat(market_data.USD_FLO_last_weighted, 'f', 5, 32)

		// fmt.Println(bittrex_share + " " + poloniex_share + " " + cryptsy_share + " " + weighted + " " + USD)
		dbtx, err := DBH.Begin()
		if err != nil {
			fmt.Println("exit 8")
			log.Fatal(err)
		}

		stmtstr := `insert into markets (unixtime, cryptsy, poloniex, bittrex, daily_volume, weighted, USD) values (strftime('%s','now'), "` + cryptsy_share + `", "` + poloniex_share + `", "` + bittrex_share + `", "` + daily_volume + `", "` + weighted + `", "` + USD + `")`

		stmt, err := dbtx.Prepare(stmtstr)
		if err != nil {
			fmt.Println("exit 9")
			log.Fatal(err)
		}

		// fmt.Printf("\nstmt: %v\n\n", stmt)

		_, stmterr := stmt.Exec()
		if err != nil {
			fmt.Println("exit 10")
			log.Fatal(stmterr)
		}

		dbtx.Commit()

		fmt.Printf(".")
		time.Sleep(20000 * time.Millisecond)
	}

}

func getConfig() API_URLs {
	// parse config
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := API_URLs{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	return configuration
}
