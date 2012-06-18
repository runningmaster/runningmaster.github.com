Ubuntu 12.04 официально поддерживает Go 1, а мужики-то не знают (с):
<pre>
sudo add-apt-repository ppa:gophers/go
sudo apt-get update && sudo apt-get install golang-stable
</pre>

Для получения версии с самыми последними изменениями следует заменить название пакета `golang-stable` на `golang-tip`.

См. также:

* [Go - UbuntuWiki](https://wiki.ubuntu.com/Go)
* [Go Language Packages (GC)](https://launchpad.net/~gophers/+archive/go)