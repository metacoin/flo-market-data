package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		fmt.Println("\nError getting Bittrex market data.")
		fmt.Printf("%v\n", err)
		return 0
	} else {
		return parse_bittrex_btc_flo_last(resp)
	}

	return 0.0
}

// get a http response from bittrex, parse it out
func parse_bittrex_btc_flo_last(resp *http.Response) float64 {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("\nError getting Bittrex market data.")
		fmt.Printf("%v\n", err)
		return 0
	} else {
		data := Bittrex_BTC_FLO{}
		err := json.Unmarshal(body, &data)
		// TODO: make this not suck
		if err != nil {
			return 0
		}

		if len(data.Result) > 0 {
			result := data.Result[0]
			return result.Last
		} else {
			fmt.Printf("\nAPI timeout (bittrex)...\n")
		}
	}

	return 0.0
}

func get_bittrex_btc_flo_volu(url string) float64 {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("\nError getting Bittrex market volume.")
		fmt.Printf("%v\n", err)
		return 0
	} else {
		return parse_bittrex_btc_flo_volu(resp)
	}

	return 0.0
}

// get a http response from bittrex, parse it out
func parse_bittrex_btc_flo_volu(resp *http.Response) float64 {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("\nError getting Bittrex market volume.")
		fmt.Printf("%v\n", err)
		return 0
	} else {
		data := Bittrex_BTC_FLO{}
		json.Unmarshal(body, &data)
		if len(data.Result) > 0 {
			result := data.Result[0]
			return result.Volume
		} else {
			fmt.Printf("\nAPI timeout (bittrex)...\n")
		}
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
		fmt.Println("\nError getting Poloniex FLO exchange rate.")
		fmt.Printf("%v\n", err)
		return 0
	} else {
		return parse_poloniex_btc_flo_last(resp)
	}

	return 0.0
}

func parse_poloniex_btc_flo_last(resp *http.Response) float64 {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("\nError getting Poloniex market data.")
		fmt.Printf("%v\n", err)
		return 0
	} else {
		var alldata []interface{}
		json.Unmarshal(body, &alldata)
		if len(alldata) > 0 {
			something := alldata[0].(map[string]interface{})
			for k, v := range something {
				if k == "rate" {
                    if v == nil {
                        return 0
                    }
					rv, err := strconv.ParseFloat(v.(string), 64)
					if err != nil {
						fmt.Println("\nCan't parse Polonirex BTC/FLO JSON.")
						fmt.Printf("%v\n", err)
						return 0
					} else {
						return rv
					}
				}
			}
		} else {
			fmt.Printf("\nAPI timeout (poloniex)\n")
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
		fmt.Println("\nError getting Poloniex market volume.")
		fmt.Printf("%v\n", err)
		return 0
	} else {
		return parse_poloniex_btc_flo_volu(resp)
	}

	return 0.0
}

func parse_poloniex_btc_flo_volu(resp *http.Response) float64 {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("\nError getting Poloniex market volume.")
		fmt.Printf("%v\n", err)
		return 0
	} else {
		var alldata interface{}
		json.Unmarshal(body, &alldata)

		// TODO: make this not suck
		if alldata == nil {
			return 0
		}

		datamap := alldata.(map[string]interface{})
		for k, v := range datamap {

			if k == "BTC_FLO" {
				vv := v.(map[string]interface{})

				for kk, vvv := range vv {

					if kk == "FLO" {
                        if vvv == nil {
                            return 0
                        }
						rv, err := strconv.ParseFloat(vvv.(string), 64)
						if err != nil {
							fmt.Println("\nError parsing Poloniex volume JSON.")
							fmt.Printf("%v\n", err)
							return 0
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
func get_cryptsy_btc_ltc_last(url string) float64 {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("\nError getting Cryptsy BTC/LTC exchange rate.")
		fmt.Printf("%v\n", err)
		return 0
	} else {
		return parse_cryptsy_btc_ltc_last(resp)
	}

	return 0.0
}

func parse_cryptsy_btc_ltc_last(resp *http.Response) float64 {

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("\nError getting Cryptsy market data.")
		fmt.Printf("%v\n", err)
		return 0
	} else {

		var alldata interface{}
		json.Unmarshal(body, &alldata)

		// TODO: make this not suck
		if alldata == nil {
			return 0
		}

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
                                        if v4 == nil {
                                            return 0.0
                                        }
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
		fmt.Println("\nError getting Cryptsy FLO/LTC exchange rate.")
		fmt.Printf("%v\n", err)
		return 0, 0
	} else {
		//fmt.Printf("getting cryptsy ltc/flo data from %v\n", url)
		return parse_cryptsy_ltc_flo_last(resp)
	}

	return 0.0, 0.0
}

// this actually returns last trade and volume
func parse_cryptsy_ltc_flo_last(resp *http.Response) (float64, float64) {

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("\nError getting Cryptsy market volume.")
		fmt.Printf("%v\n", err)
		return 0, 0
	} else {

		var alldata interface{}
		json.Unmarshal(body, &alldata)

		// TODO: make this not suck
		if alldata == nil {
			return 0, 0
		}

		something := alldata.(map[string]interface{})
		//fmt.Printf("somethingelse: %v\n", something)

		for k, v := range something {
			//fmt.Printf("k, v: %v, %v\n", k, v)
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
                                        if v4 == nil {
                                            return 0.0, 0.0
                                        }
										last, err = strconv.ParseFloat(v4.(string), 64)
										if err != nil {
											// error
											fmt.Printf("error parsing cryptsy lasttradeprice: %v\n", err)
										} else {
											// last
										}
									}
									if k4 == "volume" {
                                        if v4 == nil {
                                            return 0.0, 0.0
                                        }
										volu, err = strconv.ParseFloat(v4.(string), 64)
										if err != nil {
											// error
											fmt.Printf("error parsing cryptsy volume: %v\n", err)
										} else {
											// volume
										}
									}
								}
								//fmt.Printf("last: %v, volu: %v\n", last, volu)
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
		fmt.Println("\nError getting Bitcoinaverage market data.")
		fmt.Printf("%v\n", err)
		return 0
	} else {
		return parse_bitcoinaverage_usd(resp)
	}

	return 0.0
}

func parse_bitcoinaverage_usd(resp *http.Response) float64 {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("\nError getting Bitcoinaverage USD exchange rate.")
		fmt.Printf("%v\n", err)
		return 0
	} else {
		alldata := Bitcoinaverage{}
		err := json.Unmarshal(body, &alldata)
		// TODO: make this not suck
		if err != nil {
			return 0
		}
		return alldata.Last
	}

	return 0.0
}
