package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"ssr/public"
	"ssr/store"
)

func main() {
	log.Println("starting up!")

	router := mux.NewRouter()
	router.HandleFunc("/css/styles.css", css)
	router.HandleFunc("/favicon.ico", favicon)

	router.HandleFunc("/article/{slug}", article)
	router.HandleFunc("/about", about)
	router.HandleFunc("/", home)

	port := ":8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	log.Printf("server is listening at %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func home(w http.ResponseWriter, req *http.Request) {

	public.Home(w, public.HomeParams{
		Articles: store.ListArticles(),
	})
}

func about(w http.ResponseWriter, req *http.Request) {
	public.About(w, public.AboutParams{})
}

func article(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			//TODO: check for specifc error type or rework public.Article to return null
			w.WriteHeader(404)
			public.NotFound(w)
		}
	}()

	slug := mux.Vars(req)["slug"]
	public.Article(w, public.ArticleParams{
		Slug: slug,
	})
}

func css(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	public.Css(w)
}

func favicon(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "image/x-icon")
	public.Favicon(w)
}
