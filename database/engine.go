package database

type Engine interface {
	Open() error
	OpenRemote(host, port string) error
	Close()
	Exec(query string) (*Table, error)
}
