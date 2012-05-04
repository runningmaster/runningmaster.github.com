// Copyright 2012 Dmitriy Kovalenko (runningmaster.gmail.com). All rights reserved.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

func (p *posts) unMarkdown() {
	for _, post := range *p {
		fmt.Println(post)
	}
}

func createIndexHTML(posts *posts) {
	// process template
	err := ioutil.WriteFile(*outdir+"index.html", []byte(fmt.Sprint(*posts)), os.ModePerm)
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
	err = ioutil.WriteFile(*outdir+post.File, post.Body, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println(post.Name)
}

func goTextToBlog() {
	err := os.MkdirAll(*outdir, os.ModePerm)
	if err != nil {
		panic(err)
	}

	posts := make(posts, 0)
	posts.initFromFile("./../txt/index.json")

	for _, post := range posts {
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
