Соглашения 
----------

1. API определяется следующим общим форматом:

	```
	http://api.{domain}/{service}/{version}/{aspect}/{action}
	```

2. API реализуется только двумя HTTP методами: 

	* ``GET`` - для получения данных из сервиса
	* ``POST`` - для передачи данных в сервис

3. API предпочитает формат данных ``application/json`` (см. [JSON](http://json.org/)), в специальных же случаях может быть оговорен другой тип (см. [Content-Type](http://en.wikipedia.org/wiki/Mime_type)), например, ``text/csv`` для исторической поддержки определенного вида информации.


API 01/linkdroid
----------------

* ``POST`` ``http://api.morion.ua/01/1/data/create``

	``` json
	{
	...	
	}
	```

* ``POST`` ``http://api.morion.ua/01/1/links/create``

	``` json
	{
		"SHA": string,  // контрольная сумма 
		"NEW": number   // новое значение ключа
	}
	```
	*Присваивает контрольной сумме ссылку на новое значение эталонного ключа.*

* ``POST`` ``http://api.morion.ua/01/1/links/update``

	``` json
	{
		"OLD": number,  // предыдущее значение
		"NEW": number   // новое значение
	}
	```
	*Меняет все ссылки*

* ``POST`` ``http://api.morion.ua/01/1/links/delete``

	``` json
	{
		"OLD": number  // текущее значение
	}
	```