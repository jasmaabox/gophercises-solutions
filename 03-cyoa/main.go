package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// parseStories parses stories into a map from a JSON file
func parseStories(fname string) map[string]story {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	stories := make(map[string]story)
	json.Unmarshal(content, &stories)

	return stories
}

func main() {

	stories := parseStories("gopher.json")
	t := template.Must(template.ParseFiles("tmpl/page.html"))

	storyHandler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		arc := r.URL.EscapedPath()[1:] // this is bad

		if story, ok := stories[arc]; ok {
			t.Execute(w, story)
		} else {
			http.Redirect(w, r, "/intro", http.StatusSeeOther)
		}
	}))

	log.Println("Start server on port 8080...")
	http.ListenAndServe(":8080", storyHandler)
}
