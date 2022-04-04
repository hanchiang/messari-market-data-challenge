package models

type Metric struct {
	Market           int     `json:"market"`
	TotalVolume      float64 `json:"total_volume"`
	MeanPrice        float64 `json:"mean_price"`
	MeanVolume       float64 `json:"mean_volume"`
	Vwap             float64 `json:"volume_weighted_average_price"` // Sum(price * volume) / Sum(volume)
	PercentBuyOrders float32 `json:"percentage_buy"`
	Count            int     `json:"-"`
}
