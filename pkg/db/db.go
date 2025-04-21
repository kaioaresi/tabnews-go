package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const dbCredentials = "postgres://local_user:pwddb123@localhost:5432/tabnews?sslmode=disable"

type DBAccess interface {
	GetDBInfos() (*DbInfo, error)
}

type DBConfig struct {
	Client *sql.DB
}

type DbInfo struct {
	Version            float32 `json:"version"`
	MaxConnections     int     `json:"max_connections"`
	CurrentConnections int     `json:"current_connections"`
	Status             bool    `json:"status"`
}

func NewDBClient() (*DBConfig, error) {
	db, err := sql.Open("postgres", dbCredentials)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Error failed to connect on database - %v", err)
	}

	return &DBConfig{
		Client: db,
	}, nil
}

func (c *DBConfig) Ping() error {

	err := c.Client.Ping()
	if err != nil {
		return err
	}

	log.Println("Database connection ok!")
	return nil
}

func (c *DBConfig) getVersion() (float32, error) {
	var version float32
	err := c.Client.QueryRow("SHOW server_version;").Scan(&version)
	if err != nil {
		return 0.0, err
	}

	return version, nil
}

func (c *DBConfig) maxConnections() (int, error) {
	var maxConn int
	err := c.Client.QueryRow("show max_connections;").Scan(&maxConn)
	if err != nil {
		return 0, err
	}

	return maxConn, nil
}

func (c *DBConfig) currentConnections() (int, error) {
	var currentConnections int
	err := c.Client.QueryRow("SELECT count(*)::int FROM pg_stat_activity WHERE datname = 'tabnews';").Scan(&currentConnections)
	if err != nil {
		return 0, err
	}

	return currentConnections, nil
}

func (c *DBConfig) Close() error {
	if err := c.Client.Close(); err != nil {
		return err
	}

	return nil
}

func (c *DBConfig) GetDBInfos() (*DbInfo, error) {
	version, err := c.getVersion()
	if err != nil {
		return nil, err
	}

	currenConns, err := c.currentConnections()
	if err != nil {
		return nil, err
	}

	maxConns, err := c.maxConnections()
	if err != nil {
		return nil, err
	}

	err = c.Ping()
	if err != nil {
		return nil, err
	}

	return &DbInfo{
		Version:            version,
		MaxConnections:     maxConns,
		CurrentConnections: currenConns,
		Status:             true,
	}, nil
}
