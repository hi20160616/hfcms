package service

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms/configs"
	"google.golang.org/grpc"
)

func funHead(funName string, cfg *configs.Config) error {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recoved from %s: \n%v\n", funName, err)
		}
	}()
	if cfg == nil {
		return fmt.Errorf("%s: cfg is nil.", funName)
	}
	return nil
}

func ListArticles(ctx context.Context, in *pb.ListArticlesRequest, cfg *configs.Config) (*pb.ListArticlesResponse, error) {
	if err := funHead("ListArticles", cfg); err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(cfg.API.GRPC.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewArticlesAPIClient(conn)
	return c.ListArticles(ctx, in)
}

func GetArticle(ctx context.Context, in *pb.GetArticleRequest, cfg *configs.Config) (*pb.Article, error) {
	if err := funHead("GetArticle", cfg); err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(cfg.API.GRPC.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewArticlesAPIClient(conn)
	return c.GetArticle(ctx, in)
}

func SearchArticles(ctx context.Context, in *pb.SearchArticlesRequest, cfg *configs.Config) (*pb.SearchArticlesResponse, error) {
	if err := funHead("SearchArticles", cfg); err != nil {
		return nil, err
	}
	conn, err := grpc.Dial(cfg.API.GRPC.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewArticlesAPIClient(conn)
	return c.SearchArticles(ctx, in)
}
