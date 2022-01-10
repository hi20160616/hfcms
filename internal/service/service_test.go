package service

import (
	"context"
	"fmt"
	"testing"

	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms/configs"
)

var cfg = configs.NewConfig("hfcms")

func TestListArticles(t *testing.T) {
	in := &pb.ListArticlesRequest{}
	as, err := ListArticles(context.Background(), in, cfg)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(as.Articles)
}
