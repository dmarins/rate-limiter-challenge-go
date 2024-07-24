# Rate Limiter - FullCycle Go Expert Challenge

https://goexpert.fullcycle.com.br/pos-goexpert/

[![Go](https://img.shields.io/badge/go-1.22.4-informational?logo=go)](https://go.dev)

## Clone the project

```
$ git clone https://github.com/dmarins/rate-limiter-challenge-go.git
$ cd rate-limiter-challenge-go
```

## Download dependencies

```
$ go mod tidy
```

## Run tests

```
$ make tests
```

## Containers Up

```
$ make dc-up
```

## How to simulate the access restrictions?

```
By IP (Total of 6 requests: the first five will be released and the sixth will be restricted)

$ cd requests
$ ./testing-by-ip.sh

-----

By Token (Total of 11 requests: the first ten will be released and the eleventh will be restricted)

$ cd requests
$ ./testing-by-token.sh
```