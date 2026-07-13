package config

import "fmt"

type Postgres struct {
	DB_HOST     string `required:"true" envconfig:"DB_HOST"`
	DB_PORT     string `required:"true" envconfig:"DB_PORT"`
	DB_USER     string `required:"true" envconfig:"DB_USER"`
	DB_PASSWORD string `required:"true" envconfig:"DB_PASSWORD"`
	DB_NAME     string `required:"true" envconfig:"DB_NAME"`
	DB_SSLMODE  string `required:"true" envconfig:"DB_SSLMODE"`
}

func (pgConf *Postgres) CreateDBConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		pgConf.DB_USER,
		pgConf.DB_PASSWORD,
		pgConf.DB_HOST,
		pgConf.DB_PORT,
		pgConf.DB_NAME,
		pgConf.DB_SSLMODE,
	)
}
