package endpoint

import (
	"net/http"
	"strings"

	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
)

func Router(apis []Endpoint) *router.Router {
	router := router.New()
	for _, api := range apis {
		logrus.WithFields(logrus.Fields{
			"method":  strings.ToUpper(api.method),
			"path":    api.URL,
			"latency": api.latency,
		}).Info("Register API")
		switch strings.ToUpper(api.method) {
		case http.MethodGet:
			router.GET(api.URL, api.Handle)
		case http.MethodHead:
			router.HEAD(api.URL, api.Handle)
		case http.MethodPost:
			router.POST(api.URL, api.Handle)
		case http.MethodPut:
			router.PUT(api.URL, api.Handle)
		case http.MethodPatch:
			router.PATCH(api.URL, api.Handle)
		case http.MethodDelete:
			router.DELETE(api.URL, api.Handle)
		default:
			router.ANY(api.URL, api.Handle)
		}
	}
	return router
}
