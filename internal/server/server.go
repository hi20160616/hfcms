package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"
	"github.com/hi20160616/hfcms/internal/server/handler"
	"golang.org/x/sync/errgroup"
)

func Run(ctx context.Context, addr string, g *errgroup.Group) error {
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
	g.Go(func() error {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return ctx.Err()
	})
	g.Go(func() error {
		<-ctx.Done()
		glog.Info("Shutdown http server: %s", addr)
		return s.Shutdown(ctx)
	})
	gracefulStop(ctx, g)
	return ctx.Err()
}

func gracefulStop(ctx context.Context, g *errgroup.Group) {
	// Graceful stop
	ctx, cancel := context.WithCancel(ctx)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	g.Go(func() error {
		select {
		case sig := <-sigs:
			glog.Errorf("signal caught: %s, ready to quit...", sig.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			glog.Errorf("not canceled by context: %s", err)
		} else {
			glog.Error(err)
		}
	}
}
