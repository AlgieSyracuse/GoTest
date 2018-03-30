package main

import (
	"fmt"
	"time"
)

// Simulate ... fetch data from shared channel
//     calculation of arbitrage
func TickSimulate(tk *time.Ticker, ch chan Data) {
	rcvQ := []Data{}
	for {
		select {
		case t := <-tk.C: // calculating at time point
			SearchMinMax(rcvQ)
			fmt.Printf("%s, Simulator processes data : %v \n", t.Format("15:04:05 EST"), rcvQ)
			rcvQ = []Data{}
		case d := <-ch: // fetch data at interval
			rcvQ = append(rcvQ, d)
		}
	}
}

func SearchMinMax(rcvQ []Data) {
	var minAsk, maxBid Data
	for i, d := range rcvQ {
		if i == 0 {
			minAsk = d
			maxBid = d
		} else {
			if d.Ask < minAsk.Ask {
				minAsk = d
			}
			if d.Bid > maxBid.Bid {
				maxBid = d
			}
		}
	}
}
