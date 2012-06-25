Установка плагина [GoSublime][0] в текстовый редактор [Sublime Text 2][1] позволяет наделить последний некоторыми полезными IDE-подобными функциями, как-то:

* автозавершение кода (это делает [Gocode][2])
* переход к определению функции в исходном коде пакета (это делает [MarGo][3])
* показ документации к функции
* список объявлений
* отображение мест ошибок
* автоформатирование кода при сохранении 
* выполнение кода
* и прочая, прочая 

<pre>
$ go get -u github.com/nsf/gocode
$ go get -u github.com/DisposaBoy/MarGo
$ cd ~/.config/sublime-text-2/Packages
$ git clone git://github.com/DisposaBoy/GoSublime
</pre>

Эх, еще бы добавить в редактор `virtual space` ([обещают в 3 версии][4]), без которого мне совсем плохо (привычка, выработанная тремя пятилетками в редакторе Delphi)...

[0]: https://github.com/DisposaBoy/GoSublime
[1]: http://www.sublimetext.com/2
[2]: https://github.com/nsf/gocode
[3]: https://github.com/DisposaBoy/MarGo
[4]: http://www.sublimetext.com/forum/viewtopic.php?f=4&t=805&start=10#p25827
