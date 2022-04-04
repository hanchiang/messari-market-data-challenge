package main

import (
	"marketdata/models"
	"os"
	"testing"
)

func TestUpdateMarketMetric(t *testing.T) {
	// arrange
	var trades = []models.Trade{
		{ID: 1, Market: 1, Price: 10, Volume: 15, IsBuy: false},
		{ID: 2, Market: 1, Price: 20, Volume: 10, IsBuy: false},
		{ID: 3, Market: 2, Price: 5, Volume: 10, IsBuy: false},
		{ID: 4, Market: 1, Price: 25, Volume: 5, IsBuy: true},
		{ID: 5, Market: 2, Price: 20, Volume: 5, IsBuy: true},
	}

	var expected = []models.Metric{
		{Market: 1, TotalVolume: 15 + 10 + 5, MeanPrice: (float64(10) + 20 + 25) / 3, MeanVolume: (float64(15) + 10 + 5) / 3, Vwap: (float64(10)*15 + 20*10 + 25*5) / (15 + 10 + 5), PercentBuyOrders: float32(1) / 3, Count: 3},
		{Market: 2, TotalVolume: 10 + 5, MeanPrice: (float64(5) + 20) / 2, MeanVolume: (float64(10) + 5) / 2, Vwap: (float64(5)*10 + 20*5) / (10 + 5), PercentBuyOrders: float32(1) / 2, Count: 2},
	}

	// act
	metricByMarket := make(map[int]*models.Metric)

	for _, trade := range trades {
		updateMarketMetric(metricByMarket, &trade)
	}

	// assert
	for i := 0; i < 2; i++ {
		metric := metricByMarket[i+1]

		if metric.Market != expected[i].Market {
			t.Errorf("expected: %v, received: %v", expected[i].Market, metric.Market)
		}
		if metric.TotalVolume != expected[i].TotalVolume {
			t.Errorf("expected: %v, received: %v", expected[i].TotalVolume, metric.TotalVolume)
		}
		if metric.MeanPrice != expected[i].MeanPrice {
			t.Errorf("expected: %v, received: %v", expected[i].MeanPrice, metric.MeanPrice)
		}
		if metric.MeanVolume != expected[i].MeanVolume {
			t.Errorf("expected: %v, received: %v", expected[i].MeanVolume, metric.MeanVolume)
		}
		if metric.Vwap != expected[i].Vwap {
			t.Errorf("expected: %v, received: %v", expected[i].TotalVolume, metric.TotalVolume)
		}
		if metric.PercentBuyOrders != expected[i].PercentBuyOrders {
			t.Errorf("expected: %v, received: %v", expected[i].PercentBuyOrders, metric.PercentBuyOrders)
		}
		if metric.Count != expected[i].Count {
			t.Errorf("expected: %v, received: %v", expected[i].Count, metric.Count)
		}
	}
}

func TestReadMarketData(t *testing.T) {
	// arrange
	oldStdin := os.Stdin
	// Restore original Stdin
	defer func() {
		os.Stdin = oldStdin
	}()

	file, err := os.Open("test/resources/market_data_test.txt")
	if err != nil {
		panic(err)
	}

	os.Stdin = file

	var expected = []models.Metric{
		{Market: 1, TotalVolume: 15 + 10 + 5, MeanPrice: (float64(10) + 20 + 25) / 3, MeanVolume: (float64(15) + 10 + 5) / 3, Vwap: (float64(10)*15 + 20*10 + 25*5) / (15 + 10 + 5), PercentBuyOrders: float32(1) / 3, Count: 3},
		{Market: 2, TotalVolume: 10 + 5, MeanPrice: (float64(5) + 20) / 2, MeanVolume: (float64(10) + 5) / 2, Vwap: (float64(5)*10 + 20*5) / (10 + 5), PercentBuyOrders: float32(1) / 2, Count: 2},
	}

	// act
	metricByMarket := make(map[int]*models.Metric)
	readMarketData(metricByMarket)

	// assert
	for i := 0; i < 2; i++ {
		metric := metricByMarket[i+1]

		if metric.Market != expected[i].Market {
			t.Errorf("market %v, expected: %v, received: %v", i+1, expected[i].Market, metric.Market)
		}
		if metric.TotalVolume != expected[i].TotalVolume {
			t.Errorf("market %v, expected: %v, received: %v", i+1, expected[i].TotalVolume, metric.TotalVolume)
		}
		if metric.MeanPrice != expected[i].MeanPrice {
			t.Errorf("market %v, expected: %v, received: %v", i+1, expected[i].MeanPrice, metric.MeanPrice)
		}
		if metric.MeanVolume != expected[i].MeanVolume {
			t.Errorf("market %v, expected: %v, received: %v", i+1, expected[i].MeanVolume, metric.MeanVolume)
		}
		if metric.Vwap != expected[i].Vwap {
			t.Errorf("market %v, expected: %v, received: %v", i+1, expected[i].TotalVolume, metric.TotalVolume)
		}
		if metric.PercentBuyOrders != expected[i].PercentBuyOrders {
			t.Errorf("market %v, expected: %v, received: %v", i+1, expected[i].PercentBuyOrders, metric.PercentBuyOrders)
		}
		if metric.Count != expected[i].Count {
			t.Errorf("market %v, expected: %v, received: %v", i+1, expected[i].Count, metric.Count)
		}
	}
}
