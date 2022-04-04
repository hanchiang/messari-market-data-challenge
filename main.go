package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"marketdata/models"
	"os"
	"time"
)

func main() {
	beginTime := time.Now()

	metricByMarket := make(map[int]*models.Metric)

	readMarketData(metricByMarket)
	reportMetrics(metricByMarket)

	fmt.Println(fmt.Sprintf("Time taken: %s", time.Now().Sub(beginTime).String()))
}

func readMarketData(metricByMarket map[int]*models.Metric) {
	scanner := bufio.NewScanner(os.Stdin)

	// Discard the first "BEGIN" line
	scanner.Scan()
	scanner.Bytes()

	var bytes_read []byte
	trade := models.Trade{}

	// The provided compiled binary writes 10000001 lines instead of 10000000.
	iterations := 10000001

	count := 0
	for scanner.Scan() {
		if count == iterations {
			break
		}

		bytes_read = scanner.Bytes()

		err := json.Unmarshal(bytes_read, &trade)
		if err != nil {
			panic("Error when unmarshalling json: " + err.Error())
		}

		updateMarketMetric(metricByMarket, &trade)
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}

func updateMarketMetric(metricByMarket map[int]*models.Metric, trade *models.Trade) {
	metric, ok := metricByMarket[trade.Market]
	if !ok {
		metricByMarket[trade.Market] = &models.Metric{Market: trade.Market}
		metric = metricByMarket[trade.Market]
	}

	newCount := metric.Count + 1
	newVolume := metric.TotalVolume + trade.Volume

	// The provided binary generates ~80% of the orders to be buy orders, instead of 20%
	if trade.IsBuy {
		metric.PercentBuyOrders = (metric.PercentBuyOrders*float32(metric.Count) + 1) / float32(newCount)
	} else {
		metric.PercentBuyOrders = (metric.PercentBuyOrders * float32(metric.Count)) / float32(newCount)
	}

	metric.MeanPrice = (metric.MeanPrice*float64(metric.Count) + trade.Price) / (float64(newCount))
	metric.Vwap = (metric.Vwap*metric.TotalVolume + trade.Price*trade.Volume) / (newVolume)
	metric.TotalVolume = newVolume
	metric.MeanVolume = metric.TotalVolume / float64(newCount)

	metric.Count = newCount
}

func reportMetrics(metricByMarket map[int]*models.Metric) {
	for _, metric := range metricByMarket {
		bytes, err := json.Marshal(metric)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	}
}
