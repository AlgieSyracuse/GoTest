package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// REST api :
type REST struct {
	tk     *time.Ticker
	symbol string
	run    func(string) (Data, error)
}

func main() {
	T := time.Duration(6)
	pREST := [2]*REST{
		&REST{time.NewTicker(T * time.Second), "tBTCUSD", GetBitfinex},
		&REST{time.NewTicker(T * time.Second), "BTC-USD", GetCoinbase},
	}

	ch := make(chan Data) // shared-data channel
	defer close(ch)

	for _, pT := range pREST {
		go TickWrapper(pT, ch)
	}
	go TickSimulate(time.NewTicker(T*time.Second), ch)

	ReadCmd()
}

// TickWrapper ... rest api
/*
Bitfinix : rate limit policy 10 to 90 requests per minute
*/
func TickWrapper(pT *REST, ch chan Data) {
	defer pT.tk.Stop()
	for tic := range pT.tk.C {
		if data, err := pT.run(pT.symbol); nil == err {
			fmt.Printf("%s  (REST)get:  %+v\n", tic.Format("15:04:05 EST"), data)
			select { // non-block
			case ch <- data:
			}
		} else {
			log.Println(err)
		}
	}
}

// ReadCmd : read a line from Stdin
func ReadCmd() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n\n  ===  GO client starts... ===")
	text, _ := reader.ReadString('\n') // delimit
	fmt.Println(text)
}
