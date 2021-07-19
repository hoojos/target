package endpoint

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type APIConfig struct {
	Method             string            `yaml:"method"`
	URL                string            `yaml:"url"`
	Headers            map[string]string `yaml:"headers"`
	Body               string            `yaml:"body"`
	LatencyMilliSecond int64             `yaml:"latency_millisecond"`
}

type Config struct {
	Addr    string        `yaml:"addr"`
	Timeout time.Duration `yaml:"timeout"`
	APIs    []APIConfig   `yaml:"apis"`
}

func ParseConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, errors.WithStack(err)
	}

	return &config, nil
}

func (c Config) Endpoints() ([]Endpoint, error) {
	var err error
	endpoints := make([]Endpoint, 0, len(c.APIs))
	for _, api := range c.APIs {
		var body []byte
		if strings.HasPrefix(api.Body, "@") {
			filename := api.Body[1:]
			body, err = ioutil.ReadFile(filename)
			if err != nil {
				return nil, errors.WithMessagef(err, "read file: %s error", filename)
			}
		} else {
			body = []byte(api.Body)
		}
		api := Endpoint{
			method:  api.Method,
			URL:     api.URL,
			headers: api.Headers,
			body:    body,
			latency: time.Duration(api.LatencyMilliSecond) * time.Millisecond,
		}

		endpoints = append(endpoints, api)
	}
	return endpoints, nil
}
