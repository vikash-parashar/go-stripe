package main

import (
	"flag"
	"log"
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
	flag.StringVar(&cfg.smtp.host, "smtphost", "smtp.mailtrap.io", "smtp host")
	flag.StringVar(&cfg.smtp.username, "smtpuser", "", "smtp username")
	flag.StringVar(&cfg.smtp.password, "smtppass", "", "smtp password")
	flag.IntVar(&cfg.smtp.port, "smtpport", 587, "smtp port")

	flag.Parse()

}
