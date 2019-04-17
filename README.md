# salo-location-suggester

## Description
This is a stub service that wraps Aviasales's location suggestion API, adds custom JSON responses and simple caching

## How to try it out
`$ docker build -t salo-location-suggester`

`$ docker-compose up -d`

`$ curl -X GET 'http://127.0.0.1:8080/search?term=MOW&locale=ru'`

## What had to be done

>TODO:
 Реализовать стаб сервис ответа саловского автокомлита с ошибками и таймаутами  
 Реализовать сервис подсказок локации
 У наших друзей aviasales есть замечательный сервис подсказок, который помогает пользователям выбрать город зная всего пару букв. Но, так как у нас много виджетов, и они уже работают совсем с другим форматом, мы хотим завернуть этот сервис через себя, что бы приводить его в нужный виджетам формат. Вот пример запроса к этому сервису https://places.aviasales.ru/v2/places.json?term=Москв.. 
 Формат запроса к сервису 
 		`/search?term=mow&locale=ru`
 , формат ответа
 		`[{"slug":"MOW","subtitle":"Russia","title":"Moscow"},{"slug":"DME","subtitle":"Moscow","title":"Moscow Domodedovo Airport"}, ... ]`  
 Что мы хотим от тебя получить:
 Ссылку на репозиторий
 Инструкцию как его запустить и проверить
 
 >Немного дополнительной информации:  
 Ребята из aviasales конечно молодцы что сделали такой крутой сервис, но иногда он у них не работает или слишком долго отвечает, для нас это не приемлемо, максимальное время ответа которое может ждать виджет 3с. Хорошо что эти данные не часто обновляются, и мы сможем кэшировать результаты у себя.
 Мы любим docker и docker-compose
 Мы любим хорошие логирование, оно помогает находить проблемы
