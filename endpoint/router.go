package endpoint

import (
	"net/http"
	"strings"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"
)

func Router(apis []Endpoint) *routing.Router {
	router := routing.New()
	for _, api := range apis {
		logrus.WithFields(logrus.Fields{
			"method":  strings.ToUpper(api.method),
			"path":    api.URL,
			"latency": api.latency,
		}).Info("Register API")
		switch strings.ToUpper(api.method) {
		case http.MethodGet:
			router.Get(api.URL, api.Handle)
		case http.MethodHead:
			router.Head(api.URL, api.Handle)
		case http.MethodPost:
			router.Post(api.URL, api.Handle)
		case http.MethodPut:
			router.Put(api.URL, api.Handle)
		case http.MethodPatch:
			router.Patch(api.URL, api.Handle)
		case http.MethodDelete:
			router.Delete(api.URL, api.Handle)
		default:
			router.Any(api.URL, api.Handle)
		}
	}
	return router
}
