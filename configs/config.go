package configs

// Common - common config
type Common struct {
	ServiceHost string `env:"SERVICE_HOST" envDefault:"localhost"`
	ServicePort int    `env:"SERVICE_PORT" envDefault:"8080"`
	ConfigDB
}

// ConfigDB - database connection config
type ConfigDB struct {
	DbHost     string `env:"DB_HOST" envDefault:"localhost"`
	DbPort     int    `env:"DB_PORT" envDefault:"5432"`
	DbName     string `env:"DB_NAME" envDefault:"postgres"`
	DbUsername string `env:"DB_USERNAME" envDefault:"postgres"`
	DbPassword string `env:"DB_PASSWORD"`
	DbSslmode  string `env:"DB_SSLMODE" envDefault:"disable"`
}
