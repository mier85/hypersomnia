package server

import (
	"github.com/kstkn/hypersomnia/config"
	"github.com/kstkn/hypersomnia/handler"
	"github.com/kstkn/hypersomnia/micro"
	"github.com/micro/go-micro/client"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NewServeMux(logger handler.Logger, conf config.Config) *http.ServeMux {
	localClient := micro.NewLocalClient(
		client.NewClient(client.Registry(conf.GetRegistry())),
		conf.GetRegistry(),
		conf.RpcRequestTimeout,
	)
	webClient := micro.NewMultiWebClient(conf.Environments)
	mux := http.NewServeMux()
	mux.Handle("/", handler.NewIndexHandler(localClient, webClient))
	mux.Handle("/service", handler.NewServiceHandler(localClient, webClient))
	mux.Handle("/services", handler.NewServicesHandler(logger, localClient, webClient))
	mux.Handle("/call", handler.NewCallHandler(localClient, webClient))
	return mux
}

func StartServer(logger handler.Logger, conf config.Config) {
	log.Info("starting web server on " + conf.Addr)
	mux := NewServeMux(logger, conf)
	s := &http.Server{
		Addr:    conf.Addr,
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}
