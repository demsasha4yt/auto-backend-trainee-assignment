# auto-backend-trainee-assignment

Проект, посвященный созданию сервиса для сокращения ссылок

## Алгоритм

Алгоритм основан на представлении числа (ID ссылки в базе данных) в base62.
* При создании короткой ссылки представляем ID полной ссылки в базе данных ввиде  числа в системе счисления 62.
* При переходе по короткой ссылке получаем ID с помощью base62.Decode(path) и забираем полный URL из базы данных.

## Запуск приложения:

```
docker-compose up
```
Приложение будет доступно на порте 3000

## API

#### /api/shorten_url
* `POST` : Create a new shorten_url
Запрос:
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"url": "http://google.com"}' \
  http://localhost:3000/api/shorten_url
```
Ответ: `url` - URL сайта, `shorten_url_full` - короткая ссылка, `shorten_url` - относительня короткая ссылка
```
 {
   "url":"http://google.com",
   "shorten_url_full":"192.168.99.106:3000/{ID}",
   "shorten_url":"{ID}"
  }
```
Или код ошибки с описанием.

#### /**
Запрос:
```
curl --header "Content-Type: application/json" --request GET http://192.168.99.100:3000/R
```
Ответ: `301 Moved Permamently` или `404 Not Found`



# References
[Task definition](https://github.com/avito-tech/auto-backend-trainee-assignment)