package models

import _ "github.com/caarlos0/env"

type AppConfig struct {
	Port        string `env:"PORT,required"`
	DatabaseURL string `env:"DATABASE_URL,required"`
	Secret      string `env:"SECRET,required"`
	Key         string `env:"KEY,required"`
	Node        string `env:"NODE,required"`
	Dai         string `env:"DAI,required"`
	Symbol      string `env:"SYMBOL,required"`
	Retainer    string `env:"RETAINER,required"`
}
