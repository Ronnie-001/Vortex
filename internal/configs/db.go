package configs

import (
	"fmt"

	"github.com/apache/cassandra-gocql-driver/v2"
)

var (
	// The different hosts used for the cassandra cluster.
	ipAddrs = [...]string{"192.168.1.1", "192.168.1.2", "192.168.1.3"}
)

// DB Dependency injection; share cassandra session with other services.
type DB struct {
	dbSession *gocql.Session
}

func NewSession() (*DB, error) {
	cluster := gocql.NewCluster(ipAddrs[:]...)
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Failed to create new cassandra session: %v", err)
	}


	return &DB{dbSession: session}, nil
}

func (d *DB) Close() {
	if d.dbSession != nil {
		d.dbSession.Close()
	}
}
