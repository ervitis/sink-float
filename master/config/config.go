package config

import "github.com/kelseyhightower/envconfig"

type Server struct {
	Port int `envconfig:"PORT" default:"9098"`
}

type AppConfig struct {
	Server
}

var App AppConfig

func LoadAppConfig() {
	envconfig.MustProcess("APP", &App)
}
