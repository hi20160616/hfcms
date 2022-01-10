package main

import (
	"context"
	"flag"

	"github.com/golang/glog"
	"github.com/hi20160616/hfcms/configs"
	"github.com/hi20160616/hfcms/internal/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	cfg := configs.NewConfig("hfcms")
	glog.Info("http server start on %s", cfg.Web.Addr)
	g.Go(func() error {
		return server.Run(ctx, cfg.Web.Addr, g)
	})
}
