package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
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
	APIs []APIConfig `yaml:"apis"`
}

type API struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    []byte
	Latency time.Duration
}

func (a API) Handle(ctx *routing.Context) error {
	time.Sleep(a.Latency)
	for k, v := range a.Headers {
		ctx.Response.Header.Set(k, v)
	}
	_, err := ctx.Write(a.Body)
	return err
}

func ParseAPIs(filename string) ([]API, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var config Config
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, errors.WithStack(err)
	}

	APIs := make([]API, 0, len(config.APIs))
	for _, c := range config.APIs {
		var body []byte
		if strings.HasPrefix(c.Body, "@") {
			filename := c.Body[1:]
			body, err = ioutil.ReadFile(filename)
			if err != nil {
				return nil, errors.WithMessagef(err, "read file: %s error", filename)
			}
		} else {
			body = []byte(c.Body)
		}
		api := API{
			Method:  c.Method,
			URL:     c.URL,
			Headers: c.Headers,
			Body:    body,
			Latency: time.Duration(c.LatencyMilliSecond) * time.Millisecond,
		}

		APIs = append(APIs, api)
	}
	return APIs, nil
}

func main() {
	var (
		port    uint
		timeout time.Duration
		conf    string
	)
	flag.UintVar(&port, "port", 8080, "Listen Port")
	flag.DurationVar(&timeout, "timeout", 5*time.Second, "Read and write timeout")
	flag.StringVar(&conf, "config", "target.yaml", "Config filename")
	flag.Parse()

	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	apis, err := ParseAPIs(conf)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}

	router := routing.New()
	for _, api := range apis {
		log.Printf("%s %s %s", strings.ToUpper(api.Method), api.URL, api.Latency)
		switch strings.ToUpper(api.Method) {
		case "GET":
			router.Get(api.URL, api.Handle)
		case "POST":
			router.Post(api.URL, api.Handle)
		case "PUT":
			router.Put(api.URL, api.Handle)
		case "PATCH":
			router.Patch(api.URL, api.Handle)
		case "DELETE":
			router.Delete(api.URL, api.Handle)
		default:
			router.Any(api.URL, api.Handle)
		}
	}

	server := &fasthttp.Server{
		Handler:      router.HandleRequest,
		Name:         "Target Server",
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err = server.ListenAndServe(addr)
	if err != nil {
		log.Fatal(err)
	}
}
