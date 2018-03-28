package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// TestBitfinex : tset Bitfinex REST, display result on console
func TestBitfinex(symbol string) {
	// setup for Bitfinex
	const URI = "https://api.bitfinex.com/v2/ticker/"

	// recv data from REST api
	type Data struct {
		Title   string
		Bid     float32
		BidSize float32
		Ask     float32
		AskSize float32
	}
	resp, err := http.Get(URI + symbol)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("http_err_status : ", resp.StatusCode)
		return
	}

	var nums []float32
	if err := json.NewDecoder(resp.Body).Decode(&nums); err != nil {
		log.Println(err)
		return
	}

	data := Data{
		Title:   symbol,
		Bid:     nums[0],
		BidSize: nums[1],
		Ask:     nums[2],
		AskSize: nums[3],
	}

	fmt.Println("\n BITFINIX data: ", data)
}

func TestCoinbase(symbol string) {
	// setup for Bitfinex
	const URI = "https://api.coinbase.com/v2/prices/"

	// recv data from REST api
	type Record struct {
		Data struct {
			Base     string `json:"base"`
			Currency string `json:"currency"`
			Amount   string `json:"amount"`
		} `json:"data"`
	}
	respBuy, err1 := http.Get(URI + symbol + "/buy")
	respSell, err2 := http.Get(URI + symbol + "/sell")
	if err1 != nil || err2 != nil {
		log.Println(err1, err2)
		return
	}

	defer respBuy.Body.Close()

	if respBuy.StatusCode != http.StatusOK {
		log.Println("http_err_status : ", respBuy.StatusCode)
		return
	}
	if respSell.StatusCode != http.StatusOK {
		log.Println("http_err_status : ", respSell.StatusCode)
		return
	}

	/*
		bodyBytes, err := ioutil.ReadAll(respBuy.Body)
		fmt.Println(string(bodyBytes)) */

	var recBid, recAsk Record
	if err1, err2 := json.NewDecoder(respBuy.Body).Decode(&recBid),
		json.NewDecoder(respSell.Body).Decode(&recAsk); err1 != nil || err2 != nil {

		log.Println(err1, err2)
		return
	}

	fmt.Printf("\n Coinbase\n    symbol: %s, bid-amount: %s, ask-amoutn: %s\n",
		recBid.Data.Base+"-"+recBid.Data.Currency, recBid.Data.Amount, recAsk.Data.Amount)
}

func main() {
	TestBitfinex("tETHUSD")
	TestCoinbase("ETH-USD")
}
