package store

type Article struct {
	Title string
	Slug  string
}

var articles = []Article{
	{Title: "Predictions on the State of the Web 10 Years From Now",
		Slug: "time-capsule-predictions",
	},
	{
		Title: "Advice for Getting that First Job",
		Slug:  "advice-for-getting-that-first-job",
	},
	{
		Title: "How this Blog was Made",
		Slug:  "how-this-blog-was-made",
	},
}

func ListArticles() []Article {
	return articles
}

func GetArticle(slug string) *Article {
	for _, article := range articles {
		if article.Slug == slug {
			return &article
		}
	}
	return nil
}
