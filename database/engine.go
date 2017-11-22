package database

type Engine interface {
	Open() error
	Close()
	Exec(query string) (*Table, error)
}
