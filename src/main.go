// Copyright 2012 Dmitriy Kovalenko (runningmaster.gmail.com). All rights reserved.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	//"path/filepath"
	"github.com/russross/blackfriday"
	"time"
)

var (
	outdir *string = flag.String("out", "./../.out/", "output directory")
)

type post struct {
	Name string
	Date string
	File string
	Body []byte
	Indx int
}

type posts []*post

func (p *posts) initFromFile(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	
	err = json.Unmarshal(file, p)
	if err != nil {
		panic(err)
	}
}

func createIndexHTML(posts *posts) {
	file, err := ioutil.ReadFile("./tmp-index.html")
	if err != nil {
		panic(err)
	}

	tmp, err := template.New("index").Parse(string(file))
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmp.Execute(&buf, posts)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(*outdir+"index.html", buf.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func createPostsHTML(post *post) {
	body, err := ioutil.ReadFile("./../txt/" + post.File + ".md")
	if err != nil {
		panic(err)
	}
	post.Body = blackfriday.MarkdownCommon(body)
	// process template
	err = ioutil.WriteFile(*outdir+post.File+".html", post.Body, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println(post.Name)
}

func goTextToBlog() {
	err := os.RemoveAll(*outdir)
	if err != nil {
		panic(err)
	}
	
	err = os.MkdirAll(*outdir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	posts := make(posts, 0)
	posts.initFromFile("./../txt/index.json")

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
