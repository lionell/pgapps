# PgApps

[![Build Status](https://travis-ci.org/lionell/pgapps.svg?branch=master)](https://travis-ci.org/lionell/pgapps)

This repository contains different ways to interact with database. From simple CLI to colab notebook based on [WebSocket], deployed to the cluster.
All of these are **demos** and are not designed to be used in production.

## How it works

All the interactions with database is done via
```go
type Engine interface {
	Open() error
	OpenRemote(host, port string) error
	Close()
	Exec(query string) (*Table, error)
}
```
So that you can easily substitute database backend with anything you want(eg. [Aqua]).

## CLI

After you run the tool, you'll be welcomed with a prompt.
Just type the query you want to run and press enter. It will be executed on the backend, and you'll see the nicely formatted results.
To close the tool just enter `exit`.

```bash
$ go run cmd/cli/cli.go
> select * from users
   id     name   age
    1   Ruslan    21
    2     Dima    20
    3      Mat    18
```

## RPC

It's a client/server version that uses Remote Procedure Calls to communicate.

User interface is pretty much the same as in the [CLI](#cli) version.
To start server just run

```bash
$ go run cmd/server/server.go -port 1234
2017/11/29 03:10:30 Listening on port 1234
```

Then you can connect via client

```bash
$ go run cmd/client/client.go -host localhost -port 1234
Connection to localhost:1234 established.
> select * from users
   id     name   age
    1   Ruslan    21
    2     Dima    20
    3      Mat    18
```

## REST API

bla

## Ajax

bla

## React

bla

## WebSockets

bla

## Kubernetes

bla

## Google Cloud Platform

bla

[WebSocket]: https://en.wikipedia.org/wiki/WebSocket
[Aqua]: https://github.com/lionell/aqua
