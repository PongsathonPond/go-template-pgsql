# iDev Cms Service 🔐

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![go-version](https://img.shields.io/badge/Go-v.1.19-blue.svg)](https://gitlab.com/idev-internal/clinic-queue/clinic-appointment-service.git)
[![go-framework](https://img.shields.io/badge/Framework-gin-green.svg)](https://github.com/gin-gonic/gin)
[![godoc-ref](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org)
[![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg)](https://gitlab.com/idev-internal/clinic-queue/clinic-appointment-service.git)

template for microservice with build in authentication and authorization

## Features

- [x] ```POST``` Login
- [x] ```POST``` Logout
- [x] ```POST``` Refresh Token
- [x] ```POST``` Revoke Token
- [x] ```GET``` Profile (Me)
- [x] ```PUT``` Update Profile
- [x] ```CRUD L``` Users Management
- [x] ```CRUD L``` Roles and Permission
- [x] ```GET```  All Menus
- [x] ```Get```  All Roles

###### Remark: C=Create, R=Read, U=Update, D=Delete, L=List

## System requirements development

- [x] Docker 🐳
- [x] MongoDB
- [x] Jaeger

## Api specification document

URL: http://localhost:8080/swagger/index.html

## Installation

```
git clone https://gitlab.idevflow.com/idev-internal/idev/idev-cms-service.git
cd idev-cms-service
go mod download
```

## Pre-Require

#### Swag
converts annotations to swagger documentation
```
go install github.com/swaggo/swag/cmd/swag@latest
```
#### Mockery
provides the ability to easily generate mocks for golang interfaces
```
go install github.com/vektra/mockery/v2@latest
```

## Generate mocks
```
go generate ./...
```

## Local development
copy environment and setting value
```
cp development/local.env .env
```
generate swagger doc files
```
swag init
```

## Testing
unit testing command
```
go test ./...
```
unit testing with coverage report command
```
go test -v ./... -cover
```
integrating testing command

```
go test ./... -tags integration
```

## Run ♨
run service command
```
go run main.go
```
or run with gin live reload => [Live Reload for GO](https://medium.com/@jeyraj/live-reload-for-go-afef9c25420a)
```
gin --appPort 8080 -all -i  // first time run
gin -i
```

## Others
- Uber golang style guide [link](https://github.com/pallat/uber-go-style-guide-th).
- Practical Go: Real world advice for writing maintainable Go programs [link](https://dave.cheney.net/practical-go/presentations/qcon-china.html?fbclid=IwAR2_D2Y2HXVYUNiG3LctB0kF64YKzGUatcIHm_sLYwm9SEqEKWAd76G7NAU).
