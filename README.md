# flo-market-data

Calculates FLO market data from various markets, provies an API and pushes
updates to the block chain periodically.

## Install

flo-market-data is written in go.

It requires mattn/go-sqlite3 for database operations.

```
$ go get github.com/mattn/go-sqlite3
$ go get github.com/metacoin/flo-market-data
```

## Config

Set API to "false" if you don't want the API server to run. Otherwise, config is pretty straightforward, I think. You can find the config documentation [here][1].

## Running

Navigate to the flo-market-data directory and run the program!

Remember to include all packages:

```
$ go run *.go
```

## API

Hit this URL with a `GET` request to see the recent market data:

```
http://127.0.0.1:41290/flo-market-data/v1/getAll
```

You'll get a response like this:

```
{
    "unixtime": 1432690997,
    "cryptsy": "0.1452",
    "poloniex": "0.2948",
    "bittrex": "0.5600",
    "daily-volume": "128463.84375000",
    "weighted": "0.00000703",
    "USD": "0.00168"
}
```

## Example output

*NOTE*: API mode is enabled, verbose mode is coming soon (there is no command-line output in API mode).

If all is well, you should see something like this:

```
$ go run *.go
24hr volume: 516808.49759438

bittrex vol: 68.71 
poloniex vl: 1.25 
cryptsy vol: 30.04 

weighted   : 0.00000590
flo/USD    : 0.00142257
```

## Block chain publishing

Coming soon

# License

MIT

[1]:./docs/CONFIG.md
