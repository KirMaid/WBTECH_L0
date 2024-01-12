# WBTECH L0
Необходимо разработать демонстрационный сервис с простейшим интерфейсом, отображающий данные о заказе. Модель данных в формате JSON прилагается к заданию.

Что нужно сделать:
- Развернуть локально PostgreSQL
- - Создать свою БД
-  - Настроить своего пользователя
- - Создать таблицы для хранения полученных данных
- Разработать сервис
- - Реализовать подключение и подписку на канал в nats-streaming
- - Полученные данные записывать в БД
- - Реализовать кэширование полученных данных в сервисе (сохранять in memory)
- - В случае падения сервиса необходимо восстанавливать кэш из БД
- Запустить http-сервер и выдавать данные по id из кэша
- Разработать простейший интерфейс отображения полученных данных по id заказа

### Запуск БД и Nats-streaming
docker-compose up

### Миграция и заполнение тестовыми данными
Выполнить sql файлы в папке migrations

### Запуск клиента (Отображение заказа по uid, отправка нового заказа)
go run client.go

### Запуск сервера (Получение заказов из БД отправка на клиент, сохранение нового заказа в БД)

### Как создать новый заказ:
Отправить POST запрос на Адрес_сайта/new_order/ с данными о заказе в формате JSON.

### Как получить заказ по uid:
Отправить GET запрос на Адрес_сайта/order/{uid}. Данные отображаются в виде форматированного JSON.

### Как работает кеширование:
При старте клиента из БД берутся и кешируются первые 10 записей.
В случае, если запрашиваемый заказ не входит в эти записи, делается запрос к БД и заказ добавляется в кеш.





