Началась подготовка выпуска релиза Go 1.1, кодовая база заморожена:
[Go 1.1: Open Issues](http://swtch.com/~rsc/go11.html)

<pre>
$ go version
go version go1.0.3
</pre>

С момента выпуска релиза Go 1 тихо и незаметно вышли уже три минорных обновления:

* `go1.0.3` 24.09.2012 [что нового?](https://groups.google.com/d/msg/golang-nuts/co3SvXbGrNk/sGOmwfmBZeYJ)
* `go1.0.2` 14.06.2012 [что нового?](https://groups.google.com/forum/#!topic/golang-announce/9-f_fnXNDzw)
* `go1.0.1` 27.04.2012 [что нового?](https://groups.google.com/forum/#!topic/golang-announce/2ufDgIGFFTk)
* `go1.0.0` 28.03.2012 [release notes](http://golang.org/doc/go1.html)

Следует отметить значительный [траффик](http://code.google.com/p/go/source/list) изменений в исходном коде проекта, как по частоте коммитов так и по их авторству.

Установка из исходников вместе с экспериментальными пакетами:

<pre>
$ sudo apt-get install mercurial
$ hg clone https://code.google.com/p/go
$ cd go/src
$ ./all.bash
$ go version
go version devel +a70be086fe02 Sun Dec 16 11:51:47 2012 +0900 linux/amd64
</pre>

Мой черновик настроек:

<pre>
.profile
export GOROOT=$HOME/Development/go
export GOPATH=$HOME/Development/go-get
export PATH=$PATH:$GOROOT/bin:$GOPATH

go get -u code.google.com/p/go.talks/present
go get -u github.com/russross/blackfriday
go get -u github.com/garyburd/redigo/redis
go get -u github.com/nsf/gocode
go get -u github.com/DisposaBoy/MarGo

cd ~/.config/sublime-text-2/Packages
rm -rf GoSublime
git clone git://github.com/DisposaBoy/GoSublime
</pre>