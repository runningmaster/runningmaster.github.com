Это просто тест

``` Go

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

```