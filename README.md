# go-Middleware

## Tools

- [golang](https://go.dev/learn/) Base programming language
- [gRPC](https://grpc.io/docs/languages/go/basics/) Models and Endpoints definitions
- [gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/) REST API generator
- [Buf CLI](https://buf.build/docs/tutorials/getting-started-with-buf-cli) Protobuf package manager and build tool

## Run

```sh
go run cmd/server/main.go
```

Check ping [http://localhost:8000/v1/transaction/ping](http://localhost:8000/v1/transaction/ping)

## How To

### Define new endpoints

Endpoints can only be defined in the protocol buffer services.

The content of the folder `domain` should never be modified.
