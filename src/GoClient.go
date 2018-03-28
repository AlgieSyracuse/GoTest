package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// URIcoinbase Base URI
const URIcoinbase = "https://api.coinbase.com/v2/prices/"

// URIbitfinex  Base URI
const URIbitfinex = "https://api.bitfinex.com/v2/ticker/"

func main() {
	t := time.Duration(12)
	go TickWrapper(time.NewTicker(t*time.Second), "tBTCUSD", GetBitfinex)
	go TickWrapper(time.NewTicker(t*time.Second), "BTC-USD", GetCoinbase)

	ReadCmd()
}

/* TickWrapper ...
Bitfinix : rate limit policy 10 to 90 requests per minute
*/
func TickWrapper(tk *time.Ticker, symbol string, Do func(string) (Data, error)) {
	defer tk.Stop()

	for t := range tk.C {
		if data, err := Do(symbol); nil == err {
			fmt.Println(t.Format("Jan 2 2006 15:04:05 EST"), data)
		} else {
			log.Println(err)
		}
	}
}

// ReadCmd : read a line from Stdin
func ReadCmd() {
	// Stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Go starts...")
	text, _ := reader.ReadString('\n') // delimit
	fmt.Println(text)
}
