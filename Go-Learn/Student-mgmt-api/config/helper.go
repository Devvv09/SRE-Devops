package config

type Config struct {
	DBNAME     string `mapstructure:"DB_NAME"`
	DBHOST     string `mapstructure:"DB_HOST"`
	DBPORT     string `mapstructure:"DB_PORT"`
	SERVERADDR string `mapstructure:"SERVER_ADDR"`
	DBUSER     string `mapstructure:"DB_USER"`
	DBPASSWORD string `mapstructure:"DB_PASSWORD"`
}
