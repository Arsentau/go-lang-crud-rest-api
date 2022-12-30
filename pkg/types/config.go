package types

type Config struct {
	Port       string `mapstructure:"PORT"`
	DbHost     string `mapstructure:"POSTGRES_HOST"`
	DbUser     string `mapstructure:"POSTGRES_USER"`
	DbPort     string `mapstructure:"POSTGRES_PORT"`
	DbName     string `mapstructure:"POSTGRES_NAME"`
	DbPassword string `mapstructure:"POSTGRES_PASSWORD"`
}
