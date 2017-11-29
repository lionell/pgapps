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

bla

## RPC

bla

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
