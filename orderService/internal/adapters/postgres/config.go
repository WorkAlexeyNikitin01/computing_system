package postgres

type ConfigOrderPotgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	orderTable = "order"
)
