## Golang and MUX router REST API
---
## Start Redis Server at Default Port(6379)
```bash
redis-server
```

## Go Modules - Initialize the module for the app (from root dir)

```bash
go mod init golang-mux-api
```

## Go Modules - Download external dependencies

```bash
go mod download  
```

## Build

```bash
go build
```
 

## Test (all the tests within the controller folder)

```bash
go test golang-mux-api/controller
```

## Run

```bash
go run .
```
 

 