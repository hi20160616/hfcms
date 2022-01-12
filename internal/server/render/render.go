package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/hi20160616/hfcms/configs"
	tmpl "github.com/hi20160616/hfcms/templates"
	"github.com/yuin/goldmark"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Page struct {
	Title string
	Data  interface{}
	Cfg   *configs.Config
}

var templates = template.New("")

func init() {
	templates.Funcs(template.FuncMap{
		"summary":      summary,
		"smartTime":    smartTime,
		"smartDate":    smartDate,
		"markdown":     markdown,
		"unescapeHTML": unescapeHTML,
		"plusOne":      plusOne,
	})
	templates = template.Must(templates.ParseFS(tmpl.FS, "default/*.html"))
}

func Derive(w http.ResponseWriter, tmpl string, p *Page) {
	if err := templates.ExecuteTemplate(w, tmpl+".html", p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("err template: %s.html\n\terror: %#v", tmpl, err)
	}
}

func summary(des string) string {
	dRune := []rune(des)
	if len(dRune) <= 300 {
		return des
	}
	return string(dRune[:300])
}

func parseWithZone(t time.Time) time.Time {
	loc := time.FixedZone("UTC", 8*60*60)
	return t.In(loc)

}

func smartTime(t *timestamppb.Timestamp, site string) string {
	if site == "cna" ||
		site == "dw" ||
		site == "kabar" ||
		site == "ucpnz" ||
		site == "kyodonews" {
		return t.AsTime().Format("15:04")
	}
	return parseWithZone(t.AsTime()).Format("15:04")
}

func smartDate(t *timestamppb.Timestamp, site string) string {
	if site == "cna" ||
		site == "dw" ||
		site == "kabar" ||
		site == "ucpnz" ||
		site == "kyodonews" {
		return t.AsTime().Format("01.02")
	}
	return parseWithZone(t.AsTime()).Format("01.02")
}

func markdown(in string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(in), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func unescapeHTML(s string) template.HTML {
	return template.HTML(s)
}

func plusOne(x int) int {
	return x + 1
}
