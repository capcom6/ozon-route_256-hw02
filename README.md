# Homework 2 - Почтовый бот

Бот, позволяет агрегировать информацию с нескольких почтовых ящиков о новых сообщениях в почте. Необходимые возможности:

1. Подключение ящиков по IMAP
2. Отключение ящиков
3. Получение свежих сообщений в режиме онлайн (пулинг ящиков)

## Идея реализации

### Процесс

![Процесс](docs/images/process.jpg "Процесс")

1. Бот получает сообщения из Telegram от подписанных пользователей.
2. Пользователи имеют возможность давать команды на добавление/просмотр/удаления ящиков.
    * в рамках информации о ящике предоставляется адрес imap-сервера с портом, имя пользователя, пароль;
    * сообщение с паролем удаляется из чата *(проверить наличие такой возможности в API)*;
    * если пользователь не указал пароль, то можно спрашивать его каждый раз (но польза от такого бота сомнительна);
    * можно предлагать создать мастер-пароль, тогда при наличии нескольких ящиков достаточно будет указает его в качестве ключа расшифровки, см. п 3.
3. Информация о пользователях и их ящиках сохраняется в БД. Пароль сохраняется в зашифрованном виде. 
    * в простейшем случае ключ шифрования один;
    * чуть сложнее - индивидуальный для каждого пользователя;
    * также он может задаваться пользователем в виде мастер-пароля для одновременного доступа ко всем ящикам.
4. Пользователь может запросить новую почту.
    * бот получает список ящиков из базы;
    * расшифровывает пароли;
    * получает письма из данных ящиков;
    * возвращает пользователю список писем (либо новых в смысле "не прочитанные", либо в смысле "с прошлого запроса");
    * если писем больше некоторого количества, то выдавать самые свежие с возможность навигации по списку.
5. Пользователь может запросить удаление всех ящиков.

### Данные

![Данные](docs/images/data.jpg "Данные")

Настройки (yaml+env):

* ключ доступа к Telegram;
* параметры подключения к БД;
* ключ шифрования паролей (в зависимости от реализации);
* дополнительные параметры (значение таймаутов, адрес и порт прослушивания и т.д.).

База данных:

* пользователь (id Telegram, последняя команда);
* почтовый ящик (id владельца, адрес, зашифрованный пароль, параметры сервера);
* состояние ящика (дата/время последней проверки).

## MVP

![Состояния](docs/images/activity.jpg "Состояния")

В рамках минимальной реализации будут реализованы следующие команды в Telegram:

1. /add - добавить ящик, далее запрашивается адрес ящика, имя пользователя, пароль. Сообщение с паролем в итоге удаляется из чата.
2. /remove - отображает нумерованный список ящиков. Ввод номер ящика удаляет его. Ввод `*` удаляет все ящики. Любое другое сообщение отменяет удаление.
3. /pull - запрос новых сообщения в понятии "непрочитанные" *(требуется дополнительный анализ возможностей IMAP)*.
4. /help - справка по командам.

Пароли хранятся в БД зашифрованные единым ключом.

## Реализация

Бот состоит из двух сервисов:

1. Бэкенд - предоставляет GRPC интерфейс для выполнения операция с почтовыми ящиками и получения новых сообщений.
2. Telegram-бот - интерфейс бэкенда к Telegram, отвечает за взаимодействие с пользователями.

### Бэкенд

#### Зависимости

Для хранения данных требуется PostgreSQL. Создание необходимых таблиц может быть выполнена с помощью `goose`, файлы миграций находятся в [каталоге](assets/migrations/).

#### Настройки

Настройка возможна как через yml, так и через переменные окружения (приведены в скобках):

```yml
server:
  grpc: "localhost:8000"    # адрес и порт прослушивания для GRPC (SERVER_GRPC)
  gateway: "localhost:8001" # адрес и порт прослушивания для GRPC gateway (SERVER_GATEWAY)
database:
  host: db                  # адрес сервера PostgreSQL (DATABASE_HOST)
  port: 5432                # порт сервера PostgreSQL (DATABASE_PORT)
  database: server          # имя БД (DATABASE_NAME)
  user: server              # имя пользователя (DATABASE_USER)
  password: server_password # пароль (DATABASE_PASSWORD)
service:
  secret_key: secret        # ключ шифрования паролей (SERVICE_SECRET)
```

### Telegram-бот

Входящие сообщения получает через веб-хук. Регистрация webhook: `curl "https://api.telegram.org/bot$TELEGRAM_TOKEN/setWebhook" -H "Content-Type: application/json" -d '{"url":"https://bot.something.ru/secret_token"}'`

#### Настройки

Настройка возможна как через yml, так и через переменные окружения (приведены в скобках):

```yml
telegram:
  token: token              # ключ доступа бота (TELEGRAM_TOKEN)
http:
  listen: "localhost:3000"  # адрес и порт прослушивания обновлений (HTTP_LISTEN)
  path: path                # путь ожидания обновления (HTTP_PATH)
backend:
  host: "localhost:8000"    # адрес и порт бэкенда (BACKEND_HOST)
```

## Демонстрация

Оба сервиса были собраны в [Docker-образы](./build/package/) и размещены в рамках [реестра GitLab](https://gitlab.ozon.dev/capcom6/homework-2/container_registry) с помощью [скрипта](./scripts/docker-build.sh).

Развертывание выполнено на VPS с применением Traefik в качестве обратного прокси с автоматическим получением SSL-сертификатов от Let's Encrypt. Для этого использовался [Docker Compose](./deployments/docker-compose.yml).

Бот доступен по адресу https://t.me/route256_mail_bot

## Рабочие заметки

### Telegram

Регистрация webhook: `curl "https://api.telegram.org/bot$TELEGRAM_TOKEN/setWebhook" -H "Content-Type: application/json" -d '{"url":"https://bot.something.ru/secret_token"}'`

Удаление webhook: `curl "https://api.telegram.org/bot$TELEGRAM_TOKEN/deleteWebhook"`
