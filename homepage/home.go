package homepage

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const message = "Hello GopherCon UK 2018"

// Handlers http part handlers
type Handlers struct {
	log *logrus.Logger
}

// NewHandlers constructable function
func NewHandlers(log *logrus.Logger) *Handlers {
	return &Handlers{
		log: log,
	}
}

// SetupRoutes set up url prefix & handle function
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.logger(h.home))
}

/* middlewares */

func (h *Handlers) logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.log.Infof("request processed in %s", time.Now().Sub(startTime))
		next(w, r)
	}
}

/* handle functions */

func (h *Handlers) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
