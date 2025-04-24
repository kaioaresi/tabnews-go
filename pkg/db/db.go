package db

type DBAccess interface {
	GetDBInfos() (*DbInfo, error)
	Ping() error
	Close() error
}
