package load

import "github.com/spf13/viper"

type Postgres struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

type Config struct {
	Postgres Postgres

	DebtServiceHost string
	DebtServicePort int
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		Postgres: Postgres{
			Host:     viper.GetString("postgres.host"),
			Port:     viper.GetInt("postgres.port"),
			Database: viper.GetString("postgres.dbname"),
			User:     viper.GetString("postgres.user"),
			Password: viper.GetString("postgres.password"),
		},

		DebtServiceHost: viper.GetString("service.host"),
		DebtServicePort: viper.GetInt("service.port"),
	}

	return &cfg, nil
}
