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

## Containers Up

```
$ make dc-up
```

## Run tests

```
$ make tests
```

## How to simulate the access restrictions with Redis?

```
By IP (Total of 6 requests: the first five will be released and the sixth will be restricted)

$ cd requests
$ chmod +x testing-by-ip.sh
$ ./testing-by-ip.sh

-----

By Token (Total of 11 requests: the first ten will be released and the eleventh will be restricted)

$ cd requests
$ chmod +x testing-by-token.sh
$ ./testing-by-token.sh
```

## How to simulate the access restrictions with In-Memory?

1. execute `make dc-down` to kill all containers in execution
2. open the `.env` file
3. change the var `STRATEGY` to `in-memory`
4. execute again `make dc-up` and follow instructions bellow

```
By IP (Total of 6 requests: the first five will be released and the sixth will be restricted)

$ cd requests
$ chmod +x testing-by-ip.sh
$ ./testing-by-ip.sh

-----

By Token (Total of 11 requests: the first ten will be released and the eleventh will be restricted)

$ cd requests
$ chmod +x testing-by-token.sh
$ ./testing-by-token.sh
```