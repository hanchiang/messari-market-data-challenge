# Introduction
This is the [Messari market data coding challenge](https://engineering.messari.io/blog/messari-market-data-coding-challenge)

It does 3 things:
* Reads trade data from stdin line by line
* Compute various metrics(total volume, mean volume, volume weighted average price, percentage of buy orders) for each market, as each line of data is being read
* Output metrics to stdout

# How to use
* Install [Go](https://go.dev/doc/install)
* Download the binary which prints many trade data to stdout from the challenge link above 
* Run program: `<path/to/downloaded/binary> | go run main.go`, e.g. `./Binaries_\ stdoutinator_arm64_darwin.bin | go run main.go`

# Sample output
```bash
{"market":1282,"total_volume":2835700.9919773084,"mean_price":34.511833368708096,"mean_volume":2436.1692370939077,"volume_weighted_average_price":34.50780460159236,"percentage_buy":1}
{"market":6735,"total_volume":2844821.9756010375,"mean_price":27.499944549792072,"mean_volume":2444.0051336778674,"volume_weighted_average_price":27.498082345990188,"percentage_buy":0}
```