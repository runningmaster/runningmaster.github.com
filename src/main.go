// Copyright 2012 Dmitriy Kovalenko (runningmaster.gmail.com). All rights reserved.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

var (
	out string
	src string
	txt string
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Post struct {
	Name string
	Date string
	File string
	Body string
	Indx int
}

type Posts []*Post

func (p *Posts) initFromFile(filename string) {
	file, err := ioutil.ReadFile(filename)
	checkError(err)
	err = json.Unmarshal(file, p)
	checkError(err)
}

func applyTemplate(data interface{}) []byte {
	file, err := ioutil.ReadFile("dsgn.html")
	checkError(err)
	tmp := template.Must(template.New("index").Parse(string(file)))
	var buf bytes.Buffer
	err = tmp.Execute(&buf, data)
	checkError(err)
	return buf.Bytes()
}

func createIndexHTML(posts *Posts) {
	d := struct {
		Index bool
		Posts *Posts
	}{
		true,
		posts,
	}
	err := ioutil.WriteFile(filepath.Join(out, "index.html"), applyTemplate(&d), os.ModePerm)
	checkError(err)
}

func createPostsHTML(post *Post) {
	body, err := ioutil.ReadFile(filepath.Join(txt, post.File+".md"))
	checkError(err)
	d := struct {
		Index bool
		Post  *Post
	}{
		false,
		post,
	}
	post.Body = string(blackfriday.MarkdownCommon(body))
	err = ioutil.WriteFile(filepath.Join(out, post.File+".html"), applyTemplate(&d), os.ModePerm)
	checkError(err)
}

func goTextToBlog() {
	posts := make(Posts, 0)
	posts.initFromFile(filepath.Join(txt, "index.json"))
	l := len(posts)
	for _, post := range posts {
		post.Indx = l
		l--
		createPostsHTML(post)
	}
	createIndexHTML(&posts)
}

func main() {
	t0 := time.Now()
	flag.Parse()
	goTextToBlog()
	t1 := time.Now()
	fmt.Printf("Elapsed time %s\n", t1.Sub(t0))
}

func init() {
	src = filepath.Dir(".")
	out = filepath.Join(src, "../")
	txt = filepath.Join(out, "txt")
}
