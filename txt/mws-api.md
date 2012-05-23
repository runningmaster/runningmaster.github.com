*документ постоянно обновляется*

Соглашения 
----------

1. API определяется следующим общим url-форматом:

	```
	https://api.{domain}/{version}/{aspect}/{action}[?key=value&...]
	```

2. API реализуется двумя HTTP методами: 

	* `GET` - для получения данных из сервиса в теле ответа на запрос
	* `POST` - для передачи данных в сервис в теле запроса

	В обоих методах обязателен параметр для аутентификации:

	```
	<url>?auth=auth-sha-value
	``` 

3. API предпочитает формат данных `application/json` (см. [JSON](http://json.org/)), как для приема, так и для передачи данных. В специальных же случаях может быть оговорен другой тип (см. [Content-Type](http://en.wikipedia.org/wiki/Mime_type)), например, `text/csv` для исторической поддержки определенного вида информации.

Данные передаются по протоколу [HTTPS](http://ru.wikipedia.org/wiki/HTTPS), поддерживающий шифрование. На текущий момент времени сертификат сервера самодписанный и игнорируется клиентом.

Linkdroid API
-------------

Облачный сервис распознавания входящих данных различного типа в реальном режиме времени. Нераспознанные названия отсылает в сервис экспертной системы для последующей их привязки к значениям ключей эталонных справочников. Распознанные названия вместе с их атрибутами отсылает для последующей обработки в соотвествующие сервисы.

* **Создание аутентификации:** `auth/create` (private)

	Создает новую аутентификацию `sha` (string) <-> `new` (int64), как соответствие двух идентификаторов (пользовательского восьми символьного идентификатора доступа к сервису и его идентификатора в экспертной системе):

	```
	POST https://api.morion.ua/1/auth/create?auth=1243b7cd HTTP/1.1
	Content-Type: application/json
	
	{
		"sha": "1243b7cd",
		"new": 7775577006791949410
	}
	```

* **Обновление аутентификации:** `auth/update` (private)

	Обновляет аутентификацию `sha` (string) новым значением `new` (int64):

	```
	POST https://api.morion.ua/1/auth/update?auth=1243b7cd HTTP/1.1
	Content-Type: application/json
	
	{
		"sha": "1243b7cd",
		"new": 1949777557700679410
	}
	```

* **Удаление аутентификации:** `auth/delete` (private)

	Удаляет аутентификацию `sha` (string):

	```
	POST https://api.morion.ua/1/auth/delete?auth=1243b7cd HTTP/1.1
	Content-Type: application/json
	
	{
		"sha": "1243b7cd",
	}
	```


* **Прием данных по их типу:** `incoming/add` (public)

	Принимает данные из различных источников. Формат и логика их последующей обработки зависит от параметра `type`, который представляется в виде первых восьми символов SHA-1 кода строковых констант, определяемых разработчиками сервиса для каждого отдельного случая.

	`720fc5af` - например, чек из аптеки в `json`:

	```
	POST https://api.morion.ua/1/incoming/add?auth=1243b7cd&type=720fc5af HTTP/1.1
	Content-Type: application/json

	{
		"id": 403,
		"name": "one",
		"quant": 1,
		"price": 0.5
	}
	```

	`1a383386` - что-то еще, например прайс-лист в `csv`:

	```
	POST https://api.morion.ua/1/incoming/add?auth=1243b7cd&type=1a383386 HTTP/1.1
	Content-Type: text/csv

	blah;blah;blah;blah;blah;
	blah;blah;blah;blah;blah;
	blah;blah;blah;blah;blah;
	blah;blah;blah;blah;blah;
	blah;blah;blah;blah;blah;
	```	

	`e17370c5` - или же можно принимать`xmmo`: 

	```
	POST https://api.morion.ua/1/incoming/add?auth=1243b7cd&type=e17370c5 HTTP/1.1
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

* **Создание ссылки:** `link/create` (private)
	
	Присваивает контрольной сумме `sha` (string) ссылку на новое значение эталонного ключа `new` (Int64):

	```
	POST https://api.morion.ua/1/link/create?auth=1243b7cd HTTP/1.1
	Content-Type: application/json
	
	{
		"sha": "7c0e9591ff9b00f397718a63ae17370c566d4ca8",
		"new": 5577006791947779410
	}
	```

* **Обновление ссылки:** `link/update` (private)
	
	Обновляет для контрольной суммы `sha` (string) ссылку на новое значение эталонного ключа `new` (Int64):

	```
	POST https://api.morion.ua/1/link/update?auth=1243b7cd HTTP/1.1
	Content-Type: application/json

	{
		"sha": "15ca051d88fec9a4fd6eb39e99b1148eb2d7e3a6",
		"new": 6791947557700779410
	}
	```

* **Удаление ссылки:** `link/delete` (private)
	
	Удаляет ссылку для контрольной суммы `sha` (string):

	```
	POST https://api.morion.ua/1/link/delete?auth=1243b7cd HTTP/1.1
	Content-Type: application/json

	{
		"sha": "97c01ac379ab5c9019a2fff10de32ecb36abae53"
	}
	```

* **Переназначение ссылок:** `links/update` (private)

	Обновляет все ссылки на эталонный ключ `old` (Int64) на его новое значение `id_new` (Int64):

	```
	POST https://api.morion.ua/1/links/update?auth=1243b7cd HTTP/1.1
	Content-Type: application/json

	{
		"old": 5577006791947779410,
		"new": 4039410712379194777
	}
	```

* **Удаление ссылок:** `links/delete` (private)

	Удаляет все ссылки на эталонный ключ `old` (Int64):

	```
	POST https://api.morion.ua/1/links/delete?auth=1243b7cd HTTP/1.1
	Content-Type: application/json

	{
		"old": 4039410712379194777
	}
	```