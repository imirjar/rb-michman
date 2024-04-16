package michman

import "github.com/caarlos0/env/v10"

type MichmanConfig struct {
	Port   string `env:"PORT"`
	Secret string `env:"DB_PATH"`
}

func NewMichmanConfig() *MichmanConfig {
	conf := MichmanConfig{}
	if err := env.Parse(&conf); err != nil {
		panic(err)
	}
	return &conf
}
