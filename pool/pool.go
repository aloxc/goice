package pool

import (
	"sync"
	"time"
)

var (
	poolMux sync.RWMutex
	drivers = make(map[string]Driver)
)

type Conn interface {
	Prepare(query string) (Stmt, error)
	Close() error
}
type Driver interface {
	// Open returns a new connection to the database.
	// The name is a string in a driver-specific format.
	//
	// Open may return a cached connection (one previously
	// closed), but doing so is unnecessary; the sql package
	// maintains a pool of idle connections for efficient re-use.
	//
	// The returned connection is only used by one goroutine at a
	// time.
	Open(name string) (Conn, error)
}
type Pool struct {
	InitCap int
	// Maxcap is max connection number of the pool
	MaxCap int
	// WaitTimeout is the timeout for waiting to borrow a connection
	WaitTimeout time.Duration
	// IdleTimeout is the timeout for a connection to be alive
	IdleTimeout time.Duration
}

func (this *Pool) Exec() {

}
func Open(driverName, dataSourceName string) (*Pool, error) {
}
