// Copyright 2012 Dmitriy Kovalenko (runningmaster.gmail.com). All rights reserved.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	//"path/filepath"
	"github.com/russross/blackfriday"
	"time"
)

var (
	outdir *string = flag.String("out", "./../.out/", "output directory")
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type post struct {
	Name string
	Date string
	File string
	Body string
	Indx int
}

type posts []*post

func (p *posts) initFromFile(filename string) {
	file, err := ioutil.ReadFile(filename)
	checkError(err)
	checkError(json.Unmarshal(file, p))
}

func applyTemplate(data interface{}) []byte {
	file, err := ioutil.ReadFile("./dsgn.html")
	checkError(err)
	tmp := template.Must(template.New("index").Parse(string(file)))
	var buf bytes.Buffer
	checkError(tmp.Execute(&buf, data))
	return buf.Bytes()
}

func createIndexHTML(p *posts) {
	d := struct {
		Index bool
		Posts *posts
	}{
		true,
		p,
	}
	checkError(ioutil.WriteFile(*outdir+"index.html", applyTemplate(&d), os.ModePerm))
}

func createPostsHTML(p *post) {
	body, err := ioutil.ReadFile("./../txt/" + p.File + ".md")
	checkError(err)
	d := struct {
		Index bool
		Post  *post
	}{
		false,
		p,
	}
	p.Body = string(blackfriday.MarkdownCommon(body))
	checkError(ioutil.WriteFile(*outdir+p.File+".html", applyTemplate(&d), os.ModePerm))
}

func goTextToBlog() {
	checkError(os.RemoveAll(*outdir))
	checkError(os.MkdirAll(*outdir, os.ModePerm))
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
