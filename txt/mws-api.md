Соглашения 
----------

1. API определяется следующим общим форматом:

	```
	http://api.{domain}/{version}/{aspect}/{action}
	```

2. API реализуется двумя HTTP методами: 

	* `GET` - для получения данных из сервиса
	* `POST` - для передачи данных в сервис

3. API предпочитает формат данных `application/json` (см. [JSON](http://json.org/)), как для приема, так и для передачи данных. В специальных же случаях может быть оговорен другой тип (см. [Content-Type](http://en.wikipedia.org/wiki/Mime_type)), например, `text/csv` для исторической поддержки определенного вида информации.


Linkdroid API
-------------

Облачный сервис распознавания входящих данных различного типа в реальном режиме времени. Нераспознанные названия отсылает в сервис экспертной системы для последующей их привязки к значениям ключей эталонных справочников. Распознанные названия вместе с их атрибутами отсылает для последующей обработки в соотвествующие сервисы.

* **Прием данных по их типу:** `incoming/add`

	Принимает данные, их формат и логика последующей обработки зависит от параметра `type`, который представляется в виде первых восьми символов SHA-1 кода строковых констант, определяемых разработчиками сервиса для каждого отдельного случая.

	`720fc5af` - например, чек из аптеки в `json`:

	```
	POST http://api.morion.ua/1/incoming/add?type=720fc5af HTTP/1.1
	Content-Type: application/json

	{
		"id": 403,
		"name": ["one", "two", "three"],
		"price": [0.5, 1.25, 3.50]
	}
	```

	`1a383386` - что-то еще, например прайс-лист в `csv`:

	```
	POST http://api.morion.ua/1/incoming/add?type=1a383386 HTTP/1.1
	Content-Type: text/csv

	blah;blah;blah;blah;blah;
	blah;blah;blah;blah;blah;
	blah;blah;blah;blah;blah;
	blah;blah;blah;blah;blah;
	blah;blah;blah;blah;blah;
	```	

	`e17370c5` - или же можно принимать`xmmo`: 

	```
	POST http://api.morion.ua/1/incoming/add?type=e17370c5 HTTP/1.1
	Content-Type: text/xml

	<?xml version="1.0"?>
	<ElOrder> 
		<Document DocType="РасходнаяНакладная" Version="1" DivType=" Простые ">
			<Header> 
			</Header> 
			<Body>
				<Items>
				</Items>
			</Body> 
		</Document> 
	</ElOrder> 
	```

* **Создание ссылки:** `link/create`
	
	Присваивает контрольной сумме `sha` (string) ссылку на новое значение эталонного ключа `new` (Int64):

	```
	POST http://api.morion.ua/1/link/create HTTP/1.1
	Content-Type: application/json
	
	{
		"sha": "7c0e9591ff9b00f397718a63ae17370c566d4ca8",
		"new": 5577006791947779410
	}
	```

* **Обновление ссылки:** `link/update`
	
	Обновляет для контрольной суммы `sha` (string) ссылку на новое значение эталонного ключа `new` (Int64):

	```
	POST http://api.morion.ua/1/link/update HTTP/1.1
	Content-Type: application/json

	{
		"sha": "15ca051d88fec9a4fd6eb39e99b1148eb2d7e3a6",
		"new": 6791947557700779410
	}
	```

* **Удаление ссылки:** `link/delete`
	
	Удаляет ссылку для контрольной суммы `sha` (string):

	```
	POST http://api.morion.ua/1/link/delete HTTP/1.1
	Content-Type: application/json

	{
		"sha": "97c01ac379ab5c9019a2fff10de32ecb36abae53"
	}
	```

* **Переназначение ссылок:** `links/update`

	Обновляет все ссылки на эталонный ключ `old` (Int64) на его новое значение `new` (Int64):

	```
	POST http://api.morion.ua/1/links/update HTTP/1.1
	Content-Type: application/json

	{
		"old": 5577006791947779410,
		"new": 4039410712379194777
	}
	```

* **Удаление ссылок:** `links/delete`

	Удаляет все ссылки на эталонный ключ `old` (Int64):

	```
	POST http://api.morion.ua/1/links/delete HTTP/1.1
	Content-Type: application/json

	{
		"old": 4039410712379194777
	}
	```