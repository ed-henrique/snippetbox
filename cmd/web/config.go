package main

import "flag"

type config struct {
	addr      string
	staticDir string
	debug     bool
}

func getConfig() (cfg config) {
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.BoolVar(&cfg.debug, "debug", false, "Debug mode")

	return
}
