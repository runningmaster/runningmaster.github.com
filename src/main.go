// Copyright 2012 Dmitriy Kovalenko (runningmaster@gmail.com). All rights reserved.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

const (
	site = "runningmaster's githublog"       // Site name
	host = "http://runningmaster.github.com" // Site address
	ghub = "http://github.com/runningmaster" // Github location
	dqus = "runningmastersgithublog"         // Disqus forum shortname
	name = "Dmitriy Kovalenko"               // Author name
	mail = "runningmaster@gmail.com"         // Author mail
)

var (
	out string
	src string
	txt string
)

type Post struct {
	Name string
	Date string
	File string
	Body string
	Indx int
}

// DISQUS
// JavaScript configuration variables
// http://docs.disqus.com/help/2/
type Dqus struct {
	Name string // disqus_shortname
	Test int    // disqus_developer
	Post string // disqus_identifier
	Addr string // disqus_url
}

type Posts []*Post

func (p *Posts) initFromFile(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, p)
	if err != nil {
		panic(err)
	}
}

func applyTemplate(filename string, data interface{}) []byte {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	tmp := template.Must(template.New("index").Parse(string(file)))

	var buf bytes.Buffer
	err = tmp.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func createAtomFeed(posts *Posts) {
	d := struct {
		Site  string
		Host  string
		Name  string
		Mail  string
		Date  string
		Posts *Posts
	}{
		site,
		host,
		name,
		mail,
		time.Now().Format(time.RFC3339),
		posts,
	}

	err := ioutil.WriteFile(filepath.Join(out, "feed.xml"), applyTemplate("atom.xml", &d), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func createHomePage(posts *Posts) {
	d := struct {
		Site  string
		Host  string
		Ghub  string
		Name  string
		Mail  string
		Index bool
		Posts *Posts
	}{
		site,
		host,
		ghub,
		name,
		mail,
		true,
		posts,
	}

	err := ioutil.WriteFile(filepath.Join(out, "index.html"), applyTemplate("dsgn.html", &d), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func createPostPage(post *Post) {
	body, err := ioutil.ReadFile(filepath.Join(txt, post.File+".md"))
	if err != nil {
		panic(err)
	}

	d := struct {
		Site  string
		Host  string
		Ghub  string
		Name  string
		Mail  string
		Index bool
		Post  *Post
		Dqus  *Dqus
	}{
		site,
		host,
		ghub,
		name,
		mail,
		false,
		post,
		&Dqus{dqus, 0, post.File, fmt.Sprintf("%s/%s.html", host, post.File)},
	}

	post.Body = string(blackfriday.MarkdownCommon(body))
	// highlight(post.Body)

	err = ioutil.WriteFile(filepath.Join(out, post.File+".html"), applyTemplate("dsgn.html", &d), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func goTextToBlog() {
	posts := make(Posts, 0)
	posts.initFromFile(filepath.Join(txt, "index.json"))

	l := len(posts)
	for _, post := range posts {
		post.Indx = l
		l--
		createPostPage(post)
	}

	createAtomFeed(&posts)
	createHomePage(&posts)
}

func main() {
	t0 := time.Now()
	fmt.Println("\nElementary static blog generator\nCopyright (c) 2012 by Dmitriy Kovalenko\n")

	goTextToBlog()

	t1 := time.Now()
	fmt.Printf("Elapsed time %s\n", t1.Sub(t0))
}

func init() {
	src = filepath.Dir(".")
	out = filepath.Join(src, "../")
	txt = filepath.Join(out, "txt")
}
