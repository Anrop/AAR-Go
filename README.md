# Arma AAR API

Simple API that returns missions and events recorded by the AAR Server

## Requirements

Code is written in [Go](https://golang.org/) and uses
[Go modules](https://github.com/golang/go/wiki/Modules) for dependency
management.

## How To Use

Use `go build` to download all dependencies and compile the sources.

Start the API with the `AAR-Go` binary.
Server will be available at `$PORT`.

## Environment Variables

Environment variables can be specified in `.env` file and will be autoloaded


| Key | Required | Description |
| --- | -------- | ----------- |
| DATABASE_URL | Yes | Postgres URL to your AAR Database |
| MAX_DATABASE_CONNECTIONS` | No | 2 by default |
| PORT | No | Port that HTTP Server is bound to. 8080 by default |
