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
	for i, a := range as.Articles {
		fmt.Println(i, a.Category)
		fmt.Println(i, a.Tags)
		fmt.Println(i, a.Attributes)
	}
}

func TestGetArticle(t *testing.T) {
	id := "211229113754.21503400003"
	in := &pb.GetArticleRequest{Name: "articles/" + id}
	a, err := GetArticle(context.Background(), in, cfg)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}
