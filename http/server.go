package http

import (
	"fmt"
	"net/http"
	"platform/config"
	"platform/logging"
	"platform/pipeline"
	"sync"
)

type pipelineAdapter struct {
	pipeline.RequestPipeline
}

func (p pipelineAdapter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	p.ProcessRequest(request, writer)
}

func Serve(pl pipeline.RequestPipeline, cfg config.Configuration, logger logging.Logger) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	adapter := pipelineAdapter{RequestPipeline: pl}
	enableHttp := cfg.GetBoolDefault("http:enableHttp", true)
	if enableHttp {
		httpPort := cfg.GetIntDefault("http:port", 5001)
		logger.Debugf("Starting HTTP server on port %v", httpPort)
		wg.Add(1)
		go func() {
			err := http.ListenAndServe(fmt.Sprintf(":%v", httpPort), adapter)
			if err != nil {
				panic(err)
			}
		}()
	}
	enableHttps := cfg.GetBoolDefault("http:enableHttps", false)
	if enableHttps {
		httpsPort := cfg.GetIntDefault("http:enableHttps", 5500)
		certFile, cfok := cfg.GetString("http:httpsCert")
		keyFile, kfok := cfg.GetString("http:httpsKey")
		if cfok && kfok {
			logger.Debugf("Starting HTTPS server on port %v", httpsPort)
			wg.Add(1)
			go func() {
				err := http.ListenAndServeTLS(fmt.Sprintf(":%v", httpsPort), certFile, keyFile, adapter)
				if err != nil {
					panic(err)
				}
			}()
		} else {
			panic("HTTPS Certificate settings not found")
		}
	}
	return &wg
}
