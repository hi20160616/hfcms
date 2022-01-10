package server

import (
	"context"
	"net/http"

	"github.com/golang/glog"
	"github.com/hi20160616/hfcms/configs"
	"github.com/hi20160616/hfcms/internal/server/handler"
)

func Run(ctx context.Context, cfg *configs.Config) error {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			glog.Error(e)
		}
	}()

	s := &http.Server{
		Addr:    cfg.Web.Addr,
		Handler: handler.GetHandler(cfg),
	}
	go func() error {
		<-ctx.Done()
		glog.Infof("Shutdown http server: %s", cfg.Web.Addr)
		return s.Shutdown(ctx)
	}()

	glog.Infof("http server start on %s", cfg.Web.Addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return ctx.Err()
}
