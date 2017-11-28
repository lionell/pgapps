# PgApps

[![Build Status](https://travis-ci.org/lionell/pgapps.svg?branch=master)](https://travis-ci.org/lionell/pgapps)

This repository shows many different interfaces to interact with database. From simple CLI to [WebSocket]-based colab notebook deployed to the cluster.

## How it works

All the interactions with database is done via `database.Engine` interface.
```go
type Engine interface {
	Open() error
	OpenRemote(host, port string) error
	Close()
	Exec(query string) (*Table, error)
}
```

[WebSocket]: https://en.wikipedia.org/wiki/WebSocket
