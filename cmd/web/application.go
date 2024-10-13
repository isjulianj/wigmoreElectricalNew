package main

import (
	"log/slog"
)

type Config struct {
	HttpPort  int
	Env       string
	AssetsDir string
}

type Application struct {
	Config Config
	Logger *slog.Logger
}

func NewApplication(config Config, logger *slog.Logger) *Application {
	return &Application{
		Config: config,
		Logger: logger,
	}
}
