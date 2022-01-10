package service

import (
	"context"

	"github.com/golang/glog"
	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms/configs"
	"google.golang.org/grpc"
)

func ListArticles(ctx context.Context, in *pb.ListArticlesRequest, cfg *configs.Config) (*pb.ListArticlesResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recoved from ListArticles: \n%#v\n", err)
		}
	}()
	conn, err := grpc.Dial(cfg.API.GRPC.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewArticlesAPIClient(conn)
	return c.ListArticles(ctx, in)
}
