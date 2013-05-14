Вышел релиз Go 1.1, [скачать](https://code.google.com/p/go/downloads/list).
Новость на [OpenNet](http://www.opennet.ru/opennews/art.shtml?num=36927).

<pre>
$ go version
go version go1.1
</pre>

* `go1.1` 13.05.2013 [release notes](http://golang.org/doc/go1.1) [open issues](http://swtch.com/~rsc/go11.html)
* `go1.0.3` 24.09.2012 [что нового?](https://groups.google.com/d/msg/golang-nuts/co3SvXbGrNk/sGOmwfmBZeYJ)
* `go1.0.2` 14.06.2012 [что нового?](https://groups.google.com/forum/#!topic/golang-announce/9-f_fnXNDzw)
* `go1.0.1` 27.04.2012 [что нового?](https://groups.google.com/forum/#!topic/golang-announce/2ufDgIGFFTk)
* `go1.0.0` 28.03.2012 [release notes](http://golang.org/doc/go1.html)

[Траффик](http://code.google.com/p/go/source/list) изменений в исходном коде проекта.

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
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
</pre>