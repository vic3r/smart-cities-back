# Mi Bici Station Service

This project is related with MiBici platform, it encourages another ways to move around the city, in this case they want more people to use bikes instead of using public transport or vehicles. The headquarters are currently located in Guadalajara.

This service will provide information which will help to rebalance the bikes from different stations, in order for improving the bikes availability per station.

## Getting Started

These instructions will help you to run, deploy or test the project

### Prerequisites

What things you need to install

Golang
- Go with brew
```
brew install go
```
- From official page
```
https://golang.org/doc/install
```


### Installing

- golangci-lint with go
```
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
```
- golangci-lint with brew

```
brew install golangci/tap/golangci-lint
```
- dep
```
go get -u github.com/golang/dep/cmd/dep
```


## Running tests

From cli run
```
go test ./...
```

### And coding style tests

From cli run

```
golangci-lint run
```

## Deployment

It is automated every time it merges to `Mater` branch via CircleCI

## Authors

* **Samuel Osuna** - *Project Leader* - [Profile](https://github.com/samosunaz)
* **Estefania CP** - *Back-end Developer* - [Profile](https://github.com/Estefycp)
* **Julia Orduño** - *Full Stack Developer* - [Profile](https://github.com/juliaorduno)
* **Victor Garcia** - *DevOps* - [Profile](https://github.com/vic3r)
