package diver

type DiverConfig struct {
	Port   string `env:"PORT"`
	Secret string `env:"DB_PATH"`
	DBPath string `env:"SECRET"`
}

func NewDiverConfig() *DiverConfig {
	return &DiverConfig{}
}
