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

`cli` is a very simple [Command Line Interface](cli-wiki) for database.

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

`client` and `server` tools use [Remote Procedure Calls](rpc-wiki) to communicate.

User interface is pretty much the same as in the [CLI](#cli) version.
To start server just run

```bash
$ go run cmd/server/server.go -port 1234
2017/11/29 03:10:30 Listening on port 1234
...
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

`rest` tool exposes [REST API](rest-api-wiki) for database.

To use it, we need to start server

```bash
$ go run cmd/rest/rest.go -port 8080
2017/11/29 03:19:41 Running server on port 8080.
...
```

Server can serve queries to any table in the database.
Different tables are exposed as a different REST resources.

Here is how you can make a request

```bash
$ curl http://localhost:8080/users/1
{
  "header": [
    "id",
    "name",
    "age"
  ],
  "rows": [
    [
      "1",
      "Ruslan",
      "21"
    ]
  ]
}
```

This table shows how to make other REST calls, and whether you need to specify JSON data in request.

| Action       | Method | Pattern       | Data |
|:-------------|:-------|:--------------|:-----|
| Get all      | GET    | {table}       | No   |
| Get by ID    | GET    | {table}/{id}  | No   |
| Add new      | POST   | {table}       | Yes  |
| Update by ID | PUT    | {table}/{id}  | Yes  |
| Remove by ID | DELETE | {table}/{id}  | No   |

For example if we want to update name for a user with `id=1` we can do it this way

```bash
$ curl -X PUT --data '{"header":["name"],"rows":[["Nick"]]}' http://localhost:8080/users/1
$ curl http://localhost:8080/users/1
{
  "header": [
    "id",
    "name",
    "age"
  ],
  "rows": [
    [
      "1",
      "Nick",
      "21"
    ]
  ]
}
```

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

## License

MIT

[WebSocket]: https://en.wikipedia.org/wiki/WebSocket
[Aqua]: https://github.com/lionell/aqua
[cli-wiki]: https://en.wikipedia.org/wiki/Command-line_interface
[rpc-wiki]: https://en.wikipedia.org/wiki/Remote_procedure_call
[rest-api-wiki]: https://en.wikipedia.org/wiki/Representational_state_transfer
