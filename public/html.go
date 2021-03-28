package public

import (
	"embed"
	"fmt"
	"html/template"
	"io"

	"ssr/store"
)

//go:embed *
var files embed.FS

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file))
}

type HomeParams struct {
	Articles []store.Article
}

func Home(w io.Writer, params HomeParams) error {
	return parse("home.html").Execute(w, params)
}

type AboutParams struct {
}

func About(w io.Writer, params AboutParams) error {
	return parse("about.html").Execute(w, params)
}

type ArticleParams struct {
	Slug string
}

func Article(w io.Writer, params ArticleParams) error {
	file := fmt.Sprintf("articles/%s.html", params.Slug)
	return parse(file).Execute(w, params)
}

func Css(w io.Writer) error {
	bytes, err := files.ReadFile("styles.css")
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	return err
}

func Favicon(w io.Writer) error {
	bytes, err := files.ReadFile("favicon.ico")
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	return err
}

func NotFound(w io.Writer) error {
	return parse("notFound.html").Execute(w, struct{}{})
}
