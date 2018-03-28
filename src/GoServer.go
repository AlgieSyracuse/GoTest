package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func TestBitfinex() {
	// setup for Bitfinex

	const bfURL = "https://api.bitfinex.com/v2"
	const bfEndpoint = "Ticker"
	const bfSymbol = "tBTCUSD"

	// recv data from REST api
	type Bitfinex struct {
		BID      float32
		BID_SIZE float32
		ASK      float32
		ASK_SIZE float32
	}
	resp, err := http.Get(bfURL + "/" + bfEndpoint + "/" + bfSymbol)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(err)
		return
	}

	var nums []float32
	if err := json.NewDecoder(resp.Body).Decode(&nums); err != nil {
		log.Println(err)
		return
	}

	data := Bitfinex{
		BID:      nums[0],
		BID_SIZE: nums[1],
		ASK:      nums[2],
		ASK_SIZE: nums[3],
	}
	fmt.Println("\n BITFINIX data: ", data)

}

func main() {
	TestBitfinex()
}
