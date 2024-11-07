package config

import (
	"log"
	"os"

	"github.com/caarlos0/env/v10"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Port string `yml:"port" env:"PORT"`
	Auth string `yaml:"auth" env:"AUTH"`
}

func New() *Config {
	cfg := Config{}
	cfg.readFile()
	cfg.readEnv()
	return &cfg
}

func (cfg *Config) readFile() {
	f, err := os.Open("config/config.yml")
	if err != nil {
		log.Fatalf("error while reading file: %s", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatalf("error while decoding: %s", err)
	}
}

func (cfg *Config) readEnv() {
	c := cfg.Port
	if err := env.Parse(cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	if c != cfg.Port {
		log.Printf("Было: %s, Стало: %s", c, cfg.Port)
	}
}
