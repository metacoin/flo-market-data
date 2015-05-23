package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// BITTREX
type Bittrex_BTC_FLO struct {
	Message string `json:"message"`
	Result  []struct {
		Ask            float64 `json:"Ask"`
		BaseVolume     float64 `json:"BaseVolume"`
		Bid            float64 `json:"Bid"`
		Created        string  `json:"Created"`
		High           float64 `json:"High"`
		Last           float64 `json:"Last"`
		Low            float64 `json:"Low"`
		MarketName     string  `json:"MarketName"`
		OpenBuyOrders  int64   `json:"OpenBuyOrders"`
		OpenSellOrders int64   `json:"OpenSellOrders"`
		PrevDay        float64 `json:"PrevDay"`
		TimeStamp      string  `json:"TimeStamp"`
		Volume         float64 `json:"Volume"`
	} `json:"result"`
	Success bool `json:"success"`
}

func get_bittrex_btc_flo_last(url string) float64 {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		return parse_bittrex_btc_flo_last(resp)
	}

	return 0.0
}

// get a http response from bittrex, parse it out
func parse_bittrex_btc_flo_last(resp *http.Response) float64 {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	} else {
		data := Bittrex_BTC_FLO{}
		json.Unmarshal(body, &data)
		result := data.Result[0]
		return result.Last
	}

	return 0.0
}

func get_bittrex_btc_flo_volu(url string) float64 {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		return parse_bittrex_btc_flo_volu(resp)
	}

	return 0.0
}

// get a http response from bittrex, parse it out
func parse_bittrex_btc_flo_volu(resp *http.Response) float64 {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	} else {
		data := Bittrex_BTC_FLO{}
		json.Unmarshal(body, &data)
		result := data.Result[0]
		return result.Volume
	}

	return 0.0
}

// POLONIEX
type Poloniex_BTC_FLO_last struct {
	Amount        float64 `json:"amount,string"`
	Date          string  `json:"date"`
	GlobalTradeID int64   `json:"globalTradeID"`
	Rate          float64 `json:"rate,string"`
	Total         float64 `json:"total,string"`
	TradeID       int64   `json:"tradeID"`
	Type          string  `json:"type"`
}

func get_poloniex_btc_flo_last(url string) float64 {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		return parse_poloniex_btc_flo_last(resp)
	}

	return 0.0
}

func parse_poloniex_btc_flo_last(resp *http.Response) float64 {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	} else {
		var alldata []interface{}
		json.Unmarshal(body, &alldata)
		something := alldata[0].(map[string]interface{})
		//fmt.Printf("somethingelse: %v\n", something)
		for k, v := range something {
			if k == "rate" {
				rv, err := strconv.ParseFloat(v.(string), 64)
				if err != nil {
					log.Fatal(err)
				} else {
					return rv
				}
			}
		}
	}

	return 0.0
}

type Poloniex_BTC_FLO_volu struct {
	BTCFLO struct {
		BTC float64 `json:"BTC,string"`
		FLO float64 `json:"FLO,string"`
	} `json:"BTC_FLO"`
}

func get_poloniex_btc_flo_volu(url string) float64 {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		return parse_poloniex_btc_flo_volu(resp)
	}

	return 0.0
}

func parse_poloniex_btc_flo_volu(resp *http.Response) float64 {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	} else {
		var alldata interface{}
		json.Unmarshal(body, &alldata)
		//fmt.Printf("%v\n", alldata)
		something := alldata.(map[string]interface{})
		//fmt.Printf("%v\n", something)

		for k, v := range something {
			//fmt.Printf("k: %v, v: %v\n", k, v)

			if k == "BTC_FLO" {
				vv := v.(map[string]interface{})
				//fmt.Printf("vv(%T): %v\n", vv, vv)

				for kk, vvv := range vv {
					//fmt.Printf("k: %v, v: %v\n", kk, vvv)

					if kk == "FLO" {
						rv, err := strconv.ParseFloat(vvv.(string), 64)
						if err != nil {
							log.Fatal(err)
						} else {
							return rv
						}
					}
				}
			}
		}
	}
	return 0.0
}

// CRYPTSY

type Cryptsy_BTC_LTC_last struct {
	Amount        float64 `json:"amount,string"`
	Date          string  `json:"date"`
	GlobalTradeID int64   `json:"globalTradeID"`
	Rate          float64 `json:"rate,string"`
	Total         float64 `json:"total,string"`
	TradeID       int64   `json:"tradeID"`
	Type          string  `json:"type"`
}

func get_cryptsy_btc_ltc_last(url string) float64 {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		return parse_cryptsy_btc_ltc_last(resp)
	}

	return 0.0
}

func parse_cryptsy_btc_ltc_last(resp *http.Response) float64 {

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	} else {

		var alldata interface{}
		json.Unmarshal(body, &alldata)
		something := alldata.(map[string]interface{})
		//fmt.Printf("somethingelse: %v\n", something)

		for k, v := range something {
			if k == "return" {
				s := v.(map[string]interface{})
				for k2, v2 := range s {
					if k2 == "markets" {
						s2 := v2.(map[string]interface{})
						for k3, v3 := range s2 {
							if k3 == "LTC" {
								s3 := v3.(map[string]interface{})
								//fmt.Printf("\ns3: %v\n", s3)
								for k4, v4 := range s3 {
									if k4 == "lasttradeprice" {
										rv, err := strconv.ParseFloat(v4.(string), 64)
										if err != nil {
											return 0.0
										} else {
											return rv
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return 0.0
}

func get_cryptsy_ltc_flo_last(url string) (float64, float64) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		return parse_cryptsy_ltc_flo_last(resp)
	}

	return 0.0, 0.0
}

// this actually returns last trade and volume
func parse_cryptsy_ltc_flo_last(resp *http.Response) (float64, float64) {

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	} else {

		var alldata interface{}
		json.Unmarshal(body, &alldata)
		something := alldata.(map[string]interface{})
		//fmt.Printf("somethingelse: %v\n", something)

		for k, v := range something {
			if k == "return" {
				s := v.(map[string]interface{})
				for k2, v2 := range s {
					if k2 == "markets" {
						s2 := v2.(map[string]interface{})
						for k3, v3 := range s2 {
							if k3 == "FLO" {
								s3 := v3.(map[string]interface{})
								//fmt.Printf("\ns3: %v\n", s3)

								var volu float64 = 0.0
								var last float64 = 0.0
								for k4, v4 := range s3 {
									if k4 == "lasttradeprice" {
										var err error
										last, err = strconv.ParseFloat(v4.(string), 64)
										if err != nil {
											// error
										} else {
											// last
										}
									}
									if k4 == "volume" {
										volu, err = strconv.ParseFloat(v4.(string), 64)
										if err != nil {
											// error
										} else {
											// volume
										}
									}
								}
								return last, volu
							}
						}
					}
				}
			}
		}
	}

	return 0.0, 0.0
}

type Bitcoinaverage struct {
	_24hAvg       float64 `json:"24h_avg"`
	Ask           float64 `json:"ask"`
	Bid           float64 `json:"bid"`
	Last          float64 `json:"last"`
	Timestamp     string  `json:"timestamp"`
	VolumeBtc     float64 `json:"volume_btc"`
	VolumePercent float64 `json:"volume_percent"`
}

func get_bitcoinaverage_usd(url string) float64 {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		return parse_bitcoinaverage_usd(resp)
	}

	return 0.0
}

func parse_bitcoinaverage_usd(resp *http.Response) float64 {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	} else {
		alldata := Bitcoinaverage{}
		json.Unmarshal(body, &alldata)
		return alldata.Last
	}

	return 0.0
}
