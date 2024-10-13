package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
)

const (
	defaultPort        = 6969
	defaultEnvironment = "development"
	defaultStaticDir   = "static"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	run(logger)
}

func run(logger *slog.Logger) {
	var cfg Config

	flag.IntVar(&cfg.HttpPort, "defaultPort", defaultPort, "Api server defaultPort")
	flag.StringVar(&cfg.Env, "env", defaultEnvironment, "Environment (development|staging|production)")
	flag.StringVar(&cfg.AssetsDir, "assetDir", defaultStaticDir, "Static asset Directory")
	//dsn := flag.String("dsn", "file:./data.db", "Sqlite connection string")

	flag.Parse()

	//db, err := database.New(*dsn)
	//if err != nil {
	//	logger.Error(err.Error())
	//	os.Exit(1)
	//}
	//
	//defer db.Close()

	app := NewApplication(cfg, logger)

	mux := MuxHandler(app)

	err := app.ServeHTTP(mux)

	if err != nil {
		log.Fatal(err)
	}

	return
}
