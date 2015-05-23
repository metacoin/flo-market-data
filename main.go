package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
}

func main() {
	conf := getConfig()
	mada := MarketData{}

	/* done
	 */
	mada.bittrex_BTC_FLO_last = get_bittrex_btc_flo_last(conf.Bittrex.BTCFLOLAST)
	mada.bittrex_BTC_FLO_volu = get_bittrex_btc_flo_volu(conf.Bittrex.BTCFLOLAST)

	mada.poloniex_BTC_FLO_last = get_poloniex_btc_flo_last(conf.Poloniex.BTCFLOLAST)
	mada.poloniex_BTC_FLO_volu = get_poloniex_btc_flo_volu(conf.Poloniex.BTCFLOVOL)

	fmt.Printf("%7.8f\n", mada.poloniex_BTC_FLO_last)
	fmt.Printf("%7.8f\n", mada.poloniex_BTC_FLO_volu)
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
