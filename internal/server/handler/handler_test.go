package handler

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/hi20160616/hfcms/configs"
	"github.com/hi20160616/hfcms/internal/server/render"
)

var cfg = configs.NewConfig("hfcms")

func TestListArticlesHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost/list", nil)
	w := httptest.NewRecorder()
	p := &render.Page{Cfg: cfg, Title: "test list"}
	listArticlesHandler(w, req, p)
	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
