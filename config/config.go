package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Diver         string `env:"DIVER_ADDR"`
	Michman       string `env:"MICHMAN_ADDR"`
	Secret        string `env:"SECRET"`
	DiverTargetDB string `env:"TARGET_DB_CONN"`
}

func NewConfig() *Config {
	err := godotenv.Load() // 👈 load .env file
	if err != nil {
		log.Fatal(err)
	}

	conf := Config{}

	if err := env.Parse(&conf); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", conf)
	return &conf
}

func (c *Config) GetDiverAddr() string {
	return c.Diver
}

func (c *Config) GetMichmanAddr() string {
	return c.Michman
}

func (c *Config) GetSecret() string {
	return c.Secret
}

func (c *Config) GetDiverTargetDB() string {
	return c.DiverTargetDB
}
