package postgres

type ConfigPostgresProduct struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	productsTable = "products"
)
