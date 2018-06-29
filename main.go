package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// func index_handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "go is cool")
// 	renderTemplate(w, "index", p)
// }

var templateDir = "./template/"

var templates = template.Must(template.ParseFiles(
	templateDir+"index.html",
	templateDir+"header.html"))

// templateDir+"footer.html"))

type Page struct {
	Title   string
	Content interface{}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data *Page) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Index page handler
func index_handler(w http.ResponseWriter, r *http.Request) {
	u, _ := r.Context().Value("email").(string)
	p := &Page{
		Title: "Home",
		Content: struct {
			Email    interface{}
			LoggedIn interface{}
		}{
			template.HTML(string(u)),
			(len(u) > 0),
		},
	}
	renderTemplate(w, "index", p)
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go is fast")

}

func api_handler(w http.ResponseWriter, r *http.Request) {
	// request and parse the front page

	resp, err := http.Get("https://news.ycombinator.com")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	// define a matcher
	matcher := func(n *html.Node) bool {
		// must check for nil values
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n.Parent.Parent, "class") == "athing"
		}
		return false
	}
	// grab all articles and print them
	articles := scrape.FindAll(root, matcher)
	for i, article := range articles {
		fmt.Fprintf(w, "%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
	}
}

func main() {
	//fmt.Println("Hello World!")
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.HandleFunc("/api/", api_handler)
	http.ListenAndServe(":8000", nil)

}
