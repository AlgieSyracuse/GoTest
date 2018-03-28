package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Data : DTO
type Data struct {
	Title string
	Bid   float64
	Ask   float64
}

// GetBitfinex : tset Bitfinex REST, display result on console
// type Record struct {
//		Title   string
//		Bid     float64
//		BidSize float64
//		Ask     float64
//		AskSize float64
//	}
func GetBitfinex(symbol string) (Data, error) {
	resp, err := http.Get(URIbitfinex + symbol)
	defer resp.Body.Close()

	if err != nil {
		return Data{}, err
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("http_err_status : ", resp.StatusCode)
		return Data{}, fmt.Errorf("   bitfinix resp code: %d", resp.StatusCode)
	}

	var nums []float64
	if err := json.NewDecoder(resp.Body).Decode(&nums); err != nil {
		log.Println(err)
		return Data{}, err
	}
	return Data{Title: symbol, Bid: nums[0], Ask: nums[2]}, nil
	/*
		rec := Record{
			Title:   symbol,
			Bid:     nums[0],
			BidSize: nums[1],
			Ask:     nums[2],
			AskSize: nums[3],
		}*/
}

// GetCoinbase :
func GetCoinbase(symbol string) (Data, error) {
	// recv data from REST api
	type Record struct {
		Data struct {
			Base     string `json:"base"`
			Currency string `json:"currency"`
			Amount   string `json:"amount"`
		} `json:"data"`
	}
	respBuy, err1 := http.Get(URIcoinbase + symbol + "/buy")
	respSell, err2 := http.Get(URIcoinbase + symbol + "/sell")
	defer func() {
		respBuy.Body.Close()
		respSell.Body.Close()
	}()

	if err1 != nil {
		return Data{}, err1
	}
	if err2 != nil {
		return Data{}, err2
	}

	if respBuy.StatusCode != http.StatusOK {
		return Data{}, fmt.Errorf("  coinbase buy-resp code: %d", respBuy.StatusCode)
	}
	if respSell.StatusCode != http.StatusOK {
		return Data{}, fmt.Errorf("  coinbase sell-resp code: %d", respSell.StatusCode)
	}

	var recBid, recAsk Record
	if err1 := json.NewDecoder(respBuy.Body).Decode(&recBid); err1 != nil {
		return Data{}, err1
	}
	if err2 := json.NewDecoder(respSell.Body).Decode(&recAsk); err2 != nil {
		return Data{}, err2
	}

	// fmt.Printf("\n Coinbase\n    symbol: %s, bid-amount: %s, ask-amoutn: %s\n", recBid.Data.Base+"-"+recBid.Data.Currency, recBid.Data.Amount, recAsk.Data.Amount)
	bidPrice, _ := strconv.ParseFloat(recBid.Data.Amount, 64)
	askPrice, _ := strconv.ParseFloat(recAsk.Data.Amount, 64)

	// coinbase bug?
	if bidPrice > askPrice {
		bidPrice, askPrice = askPrice, bidPrice
	}

	return Data{
		Title: symbol,
		Bid:   bidPrice,
		Ask:   askPrice,
	}, nil
}
