package config

type Config struct {
	DatabaseURL string `mapstructure:"database_url"`
	ServerPort  string `mapstructure:"server_port"`
	JWTSecret   string `mapstructure:"jwt_secret"`
}
