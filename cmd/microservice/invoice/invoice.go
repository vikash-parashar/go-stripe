package main

import (
	"flag"
	"log"
	"os"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	smtp struct {
		host     string
		port     int
		username string
		password string
	}
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4001, "server port to listen")
	flag.StringVar(&cfg.env, "env", "development", "Application Environment {development|production|maintenance}")

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")
	infoLog := log.New(os.Stdout, "INFO : \t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR : \t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
	}

	err := app.Server()
	if err != nil {
		log.Fatal(err)
	}
}
