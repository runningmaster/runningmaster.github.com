*документ постоянно обновляется*

## Соглашения 

Данные передаются по протоколу [HTTPS](http://ru.wikipedia.org/wiki/HTTPS), поддерживающий шифрование. На текущий момент времени сертификат сервера самодписанный и игнорируется клиентом.

1. API реализуется двумя HTTP методами: 

	* `GET` - для получения данных из сервиса в теле ответа на запрос
	* `POST` - для передачи данных в сервис в теле запроса

2. API идентификация реализуется через параметры методов:

	* `auth` (string) - выдается при регистрации клиента и идентифицирует его
	* `pass` (string) - рассчитывается по одинаковому алгоритму (как на стороне клиента, так и на стороне сервера) на основе секретного ключа, уникального для каждого клиента

	```
	<authparams> => auth=public_key&pass=sha1(secret_key+post_data)
	``` 

3. API определяется следующим общим url-форматом:

	```
	https://api.{domain}/{service+version}/{aspect}/{action}[?<authparams>&key=value&...]
	```

4. API предпочитает формат данных `application/json` (см. [JSON](http://json.org/)), как для приема, так и для передачи данных. В специальных же случаях может быть оговорен другой тип (см. [Content-Type](http://en.wikipedia.org/wiki/Mime_type)), например, `text/csv` для исторической поддержки определенного вида информации.

## Linkdroid API

Облачный сервис распознавания входящих данных различного типа в реальном режиме времени. Нераспознанные названия отсылает в сервис экспертной системы для последующей их привязки к значениям ключей эталонных справочников. Распознанные названия вместе с их атрибутами отсылает для последующей обработки в соотвествующие сервисы.

### Аутентификация - auth (private)

* **Создание/обновление аутентификации:** `/ld1/auth/set`

	Создает или обновляет аутентификацию как соответствие глобального идентификатора клиента в экспертной системе `guid` (string) к паре двух хеш-ключей - публичного `pkey` (string) и секретного `skey` (string):

	```
	POST https://api.morion.ua/ld1/auth/set?auth=1243b7cd&pass=811ede49 HTTP/1.1
	Content-Type: application/json
	
	{
		"guid": "C6847488-A51B-11E1-B30F-81436188709B",
		"pkey": "ad0f7b32c41f311160db30fd2dc5f9f913f0aa41",
		"skey": "f01fd7eb1485290c10b1ac95db9710670f89bda6"
	}
	```

* **Удаление аутентификации:** `/ld1/auth/del`

	Удаляет аутентификацию `guid` (string):

	```
	POST https://api.morion.ua/ld1/auth/del?auth=1243b7cd&pass=811ede49 HTTP/1.1
	Content-Type: application/json
	
	{
		"guid": "C6847488-A51B-11E1-B30F-81436188709B"
	}
	```

### Входящие данные - data (public)

* **Прием данных по их типу:** `/ld1/data/add`

	Принимает данные из различных источников. Формат и логика их последующей обработки зависит от параметра `type`, который задается разработчиками сервиса для каждого отдельного случая.

	`type=720fc5af` - например, чек из аптеки в `json`:

	```
	POST https://api.morion.ua/ld1/data/add?auth=1243b7cd&pass=811ede49&type=720fc5af HTTP/1.1
	Content-Type: application/json

	{
		"id": 403,
		"name": "one",
		"quant": 1,
		"price": 0.5
	}
	```

	`type=1a383386` - что-то еще, например прайс-лист в `csv`:

	```
	POST https://api.morion.ua/ld1/data/add?auth=1243b7cd&pass=811ede49&type=1a383386 HTTP/1.1
	Content-Type: text/csv

	blah;blah;blah;blah;blah;
	blah;blah;blah;blah;blah;
	blah;blah;blah;blah;blah;
	blah;blah;blah;blah;blah;
	blah;blah;blah;blah;blah;
	```	

	`type=e17370c5` - или же можно принимать `xmmo`: 

	```
	POST https://api.morion.ua/ld1/data/add?auth=1243b7cd&pass=811ede49&type=e17370c5 HTTP/1.1
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

### Связь - link (private)

* **Создание/обновление связи:** `/ld1/link/set` 
	
	Устанавливает для контрольной суммы `sha` (string) ссылку на новое значение эталонного ключа `new` (Int64):

	```
	POST https://api.morion.ua/ld1/link/set?auth=1243b7cd&pass=811ede49 HTTP/1.1
	Content-Type: application/json
	
	{
		"sha": "7c0e9591ff9b00f397718a63ae17370c566d4ca8",
		"new": 5577006791947779410
	}
	```

* **Удаление связи:** `/ld1/link/del`
	
	Удаляет ссылку для контрольной суммы `sha` (string):

	```
	POST https://api.morion.ua/ld1/link/del?auth=1243b7cd&pass=811ede49 HTTP/1.1
	Content-Type: application/json

	{
		"sha": "97c01ac379ab5c9019a2fff10de32ecb36abae53"
	}
	```

### Связи - links (private)

* **Переназначение связей:** `/ld1/links/set`

	Обновляет все ссылки на эталонный ключ `old` (Int64) на его новое значение `id_new` (Int64):

	```
	POST https://api.morion.ua/ld1/links/set?auth=1243b7cd&pass=811ede49 HTTP/1.1
	Content-Type: application/json

	{
		"old": 5577006791947779410,
		"new": 4039410712379194777
	}
	```

* **Удаление связей:** `/ld1/links/del`

	Удаляет все ссылки на эталонный ключ `old` (Int64):

	```
	POST https://api.morion.ua/ld1/links/del?auth=1243b7cd&pass=811ede49 HTTP/1.1
	Content-Type: application/json

	{
		"old": 4039410712379194777
	}
	```
