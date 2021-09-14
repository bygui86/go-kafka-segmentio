package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/segmentio/stats/v4"
	"github.com/segmentio/stats/v4/prometheus"
)

var (
	customStats *CustomStats
)

/*this sample is incomplete as it does not work as expected
TODO / WARN

*/

func main() {
	fmt.Println("START segmentio stats example")

	startStats()

	incrStop := make(chan bool)
	incrTicker := startTicker(incrStop, incrTickerHandler, 500*time.Millisecond)

	printStop := make(chan bool)
	printTicker := startTicker(printStop, printTickerHandler, 3*time.Second)

	time.Sleep(10 * time.Minute)

	stopTicker(incrTicker, incrStop)
	stopTicker(printTicker, printStop)
	stopStats()

	fmt.Println("STOP segmentio stats example")
}

// STATS

type CustomStats struct {
	ID string `tag:"statsId"`

	Messages int64 `metric:"message.count" type:"counter"`
}

func (s *CustomStats) IncreaseMessages() {
	go atomic.AddInt64(&s.Messages, 1)
}

func startStats() {
	fmt.Println("START stats")

	customStats = &CustomStats{
		ID:       "test",
		Messages: 0,
	}

	stats.Register(prometheus.DefaultHandler)
	go http.ListenAndServe(":9090", prometheus.DefaultHandler)
}

func stopStats() {
	fmt.Println("STOP stats")

	stats.Flush()
}

// TICKER

func startTicker(stop chan bool, handler func(chan bool, *time.Ticker), duration time.Duration) *time.Ticker {
	// fmt.Println("START ticker")

	ticker := time.NewTicker(duration)
	go handler(stop, ticker)

	return ticker
}

func stopTicker(ticker *time.Ticker, stop chan bool) {
	// fmt.Println("STOP ticker")

	ticker.Stop()
	stop <- true
}

func incrTickerHandler(stop chan bool, ticker *time.Ticker) {
	fmt.Println("START ticker incr-handler")
	for {
		select {
		case <-stop:
			fmt.Println("Ticker incr-handler stop")
			return
		case <-ticker.C:
			customStats.IncreaseMessages()
			// stats.Incr("message.count")
		}
	}
}

func printTickerHandler(stop chan bool, ticker *time.Ticker) {
	fmt.Println("START ticker print-handler")
	for {
		select {
		case <-stop:
			fmt.Println("Ticker print-handler stop")
			return
		case <-ticker.C:
			fmt.Printf("CustomStats: %+v \n", *customStats)
		}
	}
}
