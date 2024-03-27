package main

import "flag"

type config struct {
	dsn       string
	addr      string
	staticDir string
	debug     bool
}

func getConfig() (cfg config) {
	flag.StringVar(&cfg.dsn, "dsn", "web:12345678@/snippetbox?parseTime=true", "MariaDB data source name")
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.BoolVar(&cfg.debug, "debug", false, "Debug mode")

	return
}
