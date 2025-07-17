package config

type Config struct {
	DBDriver   string
	DBName     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	JWTSecret  string
}

// TODO: read structured config from yml files
func LoadConfig() *Config {
	return &Config{
		DBDriver:  "sqlite",
		DBName:    "test.db",
		JWTSecret: "SECRET",
	}
}
