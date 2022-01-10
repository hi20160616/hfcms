package server

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/hi20160616/hfcms/internal/server/handler"
)

func Run(ctx context.Context, addr string) error {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			glog.Error(e)
		}
	}()

	s := &http.Server{
		Addr:    addr,
		Handler: handler.GetHandler(),
	}
	go func() error {
		<-ctx.Done()
		glog.Infof("Shutdown http server: %s", addr)
		return s.Shutdown(ctx)
	}()

	glog.Infof("http server start on %s", addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return ctx.Err()
}
