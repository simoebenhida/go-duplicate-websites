package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"html/template"
	"log"
	"strings"
)

type Website struct {
	link string
	root string
}

func copyWebsite(website Website) string {
	resp, err := http.Get(website.link)
	if(err != nil) {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	s := string(body[:])
	s = strings.Replace(s, `href="`, `href="`+website.root, -1)
	s = strings.Replace(s, `src="`, `src="`+website.root, -1)

	return s
}

func htmlDocHandler(w http.ResponseWriter, r *http.Request) {
	s := copyWebsite(Website{link: "https://godoc.org/golang.org/x/net/html", root: "https://godoc.org"})

	t, _ :=	template.New("fake").Parse(s)
	fmt.Println(t.Execute(w, ""))
}

func stackOverFlowHandler(w http.ResponseWriter, r *http.Request) {
	s := copyWebsite(Website{link: "https://stackoverflow.com/", root: ""})
	
	t, _ :=	template.New("fake").Parse(s)
	fmt.Println(t.Execute(w, ""))
}


func main() {

	http.HandleFunc("/doc", htmlDocHandler)
	http.HandleFunc("/stack", stackOverFlowHandler)

	http.ListenAndServe(":8080", nil)
}