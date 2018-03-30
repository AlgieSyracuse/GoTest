package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRest(t *testing.T) {
	// funcs
	foos := []func(string) (Data, error){
		GetBitfinex,
		GetCoinbase,
	}
	// params
	bars := []string{
		"tBTCUSD",
		"BTC-USD",
	}

	var data Data
	var err error
	for i, foo := range foos {
		if data, err = foo(bars[i]); err != nil {
			t.Errorf(err.Error())
		} else {
			t.Logf("rcvd data = %+v", data)
		}
	}
}

func TestChannel(t *testing.T) {
	P := fmt.Println

	ch := make(chan rune)

	go func(tk *time.Ticker) {
		var d rune
		d = 'a'
		for tic := range tk.C {
			// Rest get 'a'
			P(tic, " fetch REST ")
			select {
			case ch <- d:
				P("ID-1, send ch<- ", d)
				d++
			}
		}
	}(time.NewTicker(2 * time.Second))

	go func(tk *time.Ticker) {
		var d rune
		d = '1'
		for tic := range tk.C {
			P(tic, " fetch REST ")
			select {
			case ch <- d:
				P("ID-2, send ch<- ", d)
				d++
			}
		}
	}(time.NewTicker(2 * time.Second))

	go func(tk *time.Ticker) {
		rcvs := []rune{}
		for {
			select {
			case t := <-tk.C:
				P(t, string(rcvs), "  (next empty rcvs)")
				rcvs = []rune{}
			case d := <-ch:
				rcvs = append(rcvs, d)
				P("d:= <-ch  ", string(rcvs))
			}
		}
	}(time.NewTicker(2 * time.Second))

	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
	}
}
