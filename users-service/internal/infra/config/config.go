package config

type DatabaseConfig struct {
	Driver                  string
	Host                    string
	User                    string
	Password                string
	Name                    string
	Port                    string
	SslMode                 string
	ConnMaxLifetimeInMinute int
	MaxOpenConns            int
	MaxIdleConns            int
	MigrationsPath          string
}

type HttpConfig struct {
	Port uint
}

var DB = DatabaseConfig{
	Host:                    getEnv("DB_HOST", "localhost"),
	User:                    getEnv("DB_USER", "username"),
	Password:                getEnv("DB_PASSWORD", "password"),
	Name:                    getEnv("DB_NAME", "users"),
	Port:                    getEnv("DB_PORT", "5432"),
	SslMode:                 getEnv("DB_SSL", "disable"),
	MaxOpenConns:            getEnvInt("DB_MAX_OPEN_CONNS", 10),
	MaxIdleConns:            getEnvInt("DB_MAX_IDLE_CONNS", 1),
	ConnMaxLifetimeInMinute: getEnvInt(("DB_MAX_LIFETIME_CONN_MIN"), 3),
	MigrationsPath:          getEnv("DB_MIGRATIONS_PATH", "app/src/migrations"),
}

var Http = HttpConfig{
	Port: getEnvUint("APP_PORT", 8080),
}

var TimeZone = getEnv("TZ", "America/Santiago")
