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
	"runtime"
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

func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

type Post struct {
	Name string
	Date string
	File string
	Body string
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
	panicIfError(err)

	err = json.Unmarshal(file, p)
	panicIfError(err)
}

func applyTemplate(filename string, data interface{}) []byte {
	file, err := ioutil.ReadFile(filename)
	panicIfError(err)

	tmp := template.Must(template.New("index").Parse(string(file)))

	var buf bytes.Buffer
	err = tmp.Execute(&buf, data)
	panicIfError(err)
	return buf.Bytes()
}

func createAtomFeed(posts *Posts, c chan<- string) {
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
	panicIfError(err)
	c <- "createAtomFeed done"
}

func createHomePage(posts *Posts, c chan<- string) {
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
	panicIfError(err)
	c <- "createHomePage done"
}

func createPostPage(post *Post, c chan<- string) {
	body, err := ioutil.ReadFile(filepath.Join(txt, post.File+".md"))
	panicIfError(err)

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
	panicIfError(err)
	c <- "createPostPage done " + post.File
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time %s\n", elapsed)
}

func main() {
	defer timeTrack(time.Now())
	fmt.Println("\nElementary static blog generator\nCopyright (c) 2012 by Dmitriy Kovalenko\n")

	posts := make(Posts, 0)
	posts.initFromFile(filepath.Join(txt, "index.json"))

	c := make(chan string)
	n := 0
	for _, post := range posts {
		n++
		go createPostPage(post, c)
	}

	n++
	go createAtomFeed(&posts, c)
	n++
	go createHomePage(&posts, c)

	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	src = filepath.Dir(".")
	out = filepath.Join(src, "../")
	txt = filepath.Join(out, "txt")
}
