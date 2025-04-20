package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const dbCredentials = "postgres://local_user:pwddb123@localhost:5432/tabnews?sslmode=disable"

type DBConfig struct {
	Client *sql.DB
}

type dbInfo struct {
	Version            float32 `json:"version"`
	MaxConnetions      int     `json:"maxconnections"`
	CurrentConnections int     `json:"currentconnections"`
	Status             bool    `json:"status"`
}

func NewDBClient() *DBConfig {
	db, err := sql.Open("postgres", dbCredentials)
	if err != nil {
		panic(err)
	}

	return &DBConfig{
		Client: db,
	}
}

func (c *DBConfig) Ping() error {

	err := c.Client.Ping()
	if err != nil {
		return err
	}

	log.Println("Database connection ok!")
	return nil
}

func (c *DBConfig) GetVersion() (float32, error) {
	var version float32
	err := c.Client.QueryRow("SHOW server_version;").Scan(&version)
	if err != nil {
		return 0.0, err
	}

	return version, nil
}

func (c *DBConfig) MaxConnetions() (int, error) {
	var MaxConnetions int
	err := c.Client.QueryRow("show max_connections;").Scan(&MaxConnetions)
	if err != nil {
		return 0, err
	}

	return MaxConnetions, nil
}

func (c *DBConfig) CurrentConnections() (int, error) {
	var CurrentConnections int
	err := c.Client.QueryRow("SELECT count(*)::int FROM pg_stat_activity WHERE datname = 'tabnews';").Scan(&CurrentConnections)
	if err != nil {
		return 0, err
	}

	return CurrentConnections, nil
}

func (c *DBConfig) Close() error {
	if err := c.Client.Close(); err != nil {
		return err
	}

	return nil
}

func (c *DBConfig) GetDBInfos() (*dbInfo, error) {
	version, err := c.GetVersion()
	if err != nil {
		return nil, err
	}

	currenConnections, err := c.CurrentConnections()
	if err != nil {
		return nil, err
	}

	maxConnetions, err := c.MaxConnetions()
	if err != nil {
		return nil, err
	}

	err = c.Ping()
	if err != nil {
		return nil, err
	}

	return &dbInfo{
		Version:            version,
		MaxConnetions:      maxConnetions,
		CurrentConnections: currenConnections,
		Status:             true,
	}, nil
}
