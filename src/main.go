// Copyright 2012 Dmitriy Kovalenko (runningmaster@gmail.com). All rights reserved.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"math/rand"
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

var zenPython = [17]string{
	"Красивое лучше, чем уродливое",
	"Явное лучше, чем неявное",
	"Простое лучше, чем сложное",
	"Сложное лучше, чем запутанное",
	"Плоское лучше, чем вложенное",
	"Разреженное лучше, чем плотное",
	"Читаемость имеет значение",
	"Особые случаи не настолько особые, чтобы нарушать правила",
	"Практичность важнее безупречности",
	"Ошибки никогда не должны замалчиваться. Если не замалчиваются явно",
	"Встретив двусмысленность, отбрось искушение угадать",
	"Должен существовать один — и, желательно, только один — очевидный способ сделать это",
	"Сейчас лучше, чем никогда. Хотя никогда зачастую лучше, чем прямо сейчас",
	"Если реализацию сложно объяснить — идея плоха",
	"Если реализацию легко объяснить — идея, возможно, хороша",
	"Пространства имён — отличная штука! Будем делать их побольше!",
	"Programming, Motherfucker",
}

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

type Posts []*Post

// DISQUS
// JavaScript configuration variables
// http://docs.disqus.com/help/2/
type Dqus struct {
	Name string // disqus_shortname
	Test int    // disqus_developer
	Post string // disqus_identifier
	Addr string // disqus_url
}

type templFeed struct {
	Site  string
	Host  string
	Name  string
	Mail  string
	Date  string
	Posts *Posts
}

type templPage struct {
	Home  string
	Site  string
	Host  string
	Ghub  string
	Name  string
	Mail  string
	Post  *Post  // for post page only
	Dqus  *Dqus  // for post page only
	Posts *Posts // for home page only
}

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
	tpl := new(templFeed)
	tpl.Site = site
	tpl.Host = host
	tpl.Name = name
	tpl.Mail = mail
	tpl.Date = time.Now().Format(time.RFC3339)
	tpl.Posts = posts

	err := ioutil.WriteFile(filepath.Join(out, "feed.xml"), applyTemplate("atom.xml", &tpl), os.ModePerm)
	panicIfError(err)
	c <- "createAtomFeed done"
}

func createHomePage(posts *Posts, c chan<- string) {
	tpl := new(templPage)
	tpl.Home = fmt.Sprintf("%s (c)", zenPython[rand.Intn(len(zenPython))])
	tpl.Site = site
	tpl.Host = host
	tpl.Ghub = ghub
	tpl.Name = name
	tpl.Mail = mail
	tpl.Posts = posts

	err := ioutil.WriteFile(filepath.Join(out, "index.html"), applyTemplate("dsgn.html", &tpl), os.ModePerm)
	panicIfError(err)
	c <- "createHomePage done"
}

func createPostPage(post *Post, c chan<- string) {
	body, err := ioutil.ReadFile(filepath.Join(txt, post.File+".md"))
	panicIfError(err)

	tpl := new(templPage)
	tpl.Site = site
	tpl.Host = host
	tpl.Ghub = ghub
	tpl.Name = name
	tpl.Mail = mail
	tpl.Post = post
	tpl.Dqus = &Dqus{dqus, 0, post.File, fmt.Sprintf("%s/%s.html", host, post.File)}

	post.Body = string(blackfriday.MarkdownCommon(body))
	// highlight(post.Body)

	err = ioutil.WriteFile(filepath.Join(out, post.File+".html"), applyTemplate("dsgn.html", &tpl), os.ModePerm)
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
	rand.Seed(time.Now().UnixNano())

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
