package gohealthcheck

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type healthCheck struct {
	logger *log.Logger
	exit   func(code int)
	run    bool
}

func (hc *healthCheck) register(addr string) {
	if !hc.run {
		return
	}
	resp, err := http.Get(addr)
	if err != nil || resp.StatusCode != 200 {
		hc.logger.Println("FAIL")
		hc.exit(1)
	}
	hc.logger.Println("OK")
	hc.exit(0)
}

func defaultHc() *healthCheck {
	f := flag.Bool("hc", false, "perform a health check")
	flag.Parse()
	return &healthCheck{
		logger: log.New(os.Stdout, "health-check ", log.LstdFlags),
		exit:   os.Exit,
		run:    *f,
	}
}

func Register(addr string) {
	defaultHc().register(addr)
}
