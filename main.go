package main

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"target/endpoint"
)


func main() {
	var conf    string
	flag.StringVar(&conf, "config", "target.yaml", "Config filename")
	flag.Parse()

	config, err := endpoint.ParseConfig(conf)
	if err != nil {
		logrus.WithError(err).Error("parse config error")
	}

	endpoints, err := config.Endpoints()
	if err != nil {
		logrus.WithError(err).Error("parse endpoints error")
	}
	router := endpoint.Router(endpoints)
	server := &fasthttp.Server{
		Handler:      router.HandleRequest,
		Name:         "Target Server",
		ReadTimeout:  config.Timeout,
		WriteTimeout: config.Timeout,
	}
	err = server.ListenAndServe(config.Addr)
	if err != nil {
		logrus.WithError(err).Error("Listen error")
	}
}
