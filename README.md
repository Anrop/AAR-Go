# Arma AAR API

Simple API that returns missions and events recorded by the AAR Server

## Requirements

Code is written in [Go](https://golang.org/) and uses [gb](https://getgb.io/) to compile

## How To Use

Compile the sources with `gb build`

Environment variables can be specified in `.env` file and will be autoloaded

Set `DATABASE_URL` as a postgres url to your AAR DATABASE_URL

`MAX_DATABASE_CONNECTIONS` can be defined and is 2 by default

`PORT` can be defined and is 8080 by default

Start the API with the aar binary in the bin folder
