package internal

import (
	"context"
	"fmt"
	"github.com/lingyao2333/mo-zero/core/logx"
	"github.com/lingyao2333/mo-zero/core/proc"
	"net"
)

// StartOption defines the method to customize http.Server.
type StartOption func(svr *http.Server)

// StartHttp starts a http server.
func StartHttp(host string, port int, handler http.Handler, opts ...StartOption) error {
	return start(host, port, handler, func(svr *http.Server) error {
		//return svr.ListenAndServe()
		l, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			panic("failed to listen:" + err.Error())
		}
		return svr.Serve(l)
	}, opts...)
}

// StartHttps starts a https server.
func StartHttps(host string, port int, certFile, keyFile string, handler http.Handler,
	opts ...StartOption) error {
	return start(host, port, handler, func(svr *http.Server) error {
		// certFile and keyFile are set in buildHttpsServer
		return svr.ListenAndServeTLS(certFile, keyFile)
	}, opts...)
}

func start(host string, port int, handler http.Handler, run func(svr *http.Server) error,
	opts ...StartOption) (err error) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: handler,
	}
	for _, opt := range opts {
		opt(server)
	}

	waitForCalled := proc.AddWrapUpListener(func() {
		if e := server.Shutdown(context.Background()); e != nil {
			logx.Error(e)
		}
	})
	defer func() {
		if err == http.ErrServerClosed {
			waitForCalled()
		}
	}()

	return run(server)
}
