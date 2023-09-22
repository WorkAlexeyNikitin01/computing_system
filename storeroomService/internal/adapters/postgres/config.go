package postgres

type ConfigPostgresStoreroom struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	storeroomTable = "storeroom"
)
