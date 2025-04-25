package db

type DBAccess interface {
	GetDBInfos() (*DbInfo, error)
	Close() error
}
