package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/wotmshuaisi/gopher2018microserivceexample/homepage"
	"github.com/wotmshuaisi/gopher2018microserivceexample/server"
	"github.com/wotmshuaisi/gopher2018microserivceexample/utils"
)

var (
	certFile    = utils.GetSysEnv("EXAMPLE_CERT_FILE", "certs/example.crt")
	keyFile     = utils.GetSysEnv("EXAMPLE_KEY_FILE", "certs/example.key")
	serviceAddr = utils.GetSysEnv("EXAMPLE_SERVICE_ADDR", "0.0.0.0:8080")
)

func main() {
	utils.InitLogger()
	weblogger := utils.NewWebLogger()

	h := homepage.NewHandlers(weblogger)

	mux := http.NewServeMux()

	h.SetupRoutes(mux)

	srv := server.New(mux, serviceAddr)

	logrus.Infof("=> server listening on %v", serviceAddr)
	err := srv.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		logrus.Fatalf("server failed to start: %v", err)
	}

}
