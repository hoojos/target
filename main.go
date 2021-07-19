package main

import (
	"flag"
	"github.com/hoojos/target/endpoint"
	"github.com/hoojos/target/log"
	"github.com/hoojos/target/middleware"
	"github.com/hoojos/target/middleware/logging"
	"github.com/hoojos/target/middleware/recovery"
	"github.com/valyala/fasthttp"
)

func main() {
	var conf string
	flag.StringVar(&conf, "config", "target.yaml", "Config filename")
	flag.Parse()

	config, err := endpoint.ParseConfig(conf)
	if err != nil {
		log.WithError(err).Error("parse config error")
	}

	endpoints, err := config.Endpoints()
	if err != nil {
		log.WithError(err).Error("parse endpoints error")
	}
	router := endpoint.Router(endpoints)
	middlewares := middleware.Chain(logging.Use(log.DefaultLogger), recovery.Recovery())
	server := &fasthttp.Server{
		Handler:      middlewares(router.Handler),
		Name:         "Target Server",
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
	}
	err = server.ListenAndServe(config.Addr)
	if err != nil {
		log.WithError(err).Error("Listen error")
	}
}
