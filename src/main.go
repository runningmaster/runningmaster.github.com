// Copyright 2012 Dmitriy Kovalenko runningmaster.gmail.com. All rights reserved.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Post struct {
	Name string
	Date string
	File string
	Body string
}

type Posts []Post

func (p *Posts) InitFromFile(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, p)
	if err != nil {
		panic(err)
	}
}

func (p *Posts) UnMarkdown() {
	for _, post := range *p {
		fmt.Println(post)
	}
}

func (p *Posts) SaveSiteHTML() {
	for _, post := range *p {
		CreatePostsHTML(post)
	}
	CreateIndexHTML(p)
}

func CreateIndexHTML(posts *Posts) {
	// process template
	err := ioutil.WriteFile("./../.out/"+"index.html", []byte(fmt.Sprint(posts)), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func CreatePostsHTML(post Post) {
	// unmarkdown
	// process template
	err := ioutil.WriteFile("./../.out/"+post.File, []byte(post.Body), os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println(post.Name)
}

func GoTextToBlog() {
	posts := make(Posts, 1)
	posts.InitFromFile("./../txt/index.json")
	posts.SaveSiteHTML()
}

func main() {
	t0 := time.Now()
	err := os.RemoveAll("./../.out/")
	if err != nil {
		panic(err)
	}
	err = os.Mkdir("./../.out/", os.ModePerm)
	if err != nil {
		panic(err)
	}
	GoTextToBlog()
	t1 := time.Now()
	fmt.Printf("Elapsed time %s\n", t1.Sub(t0))
}
