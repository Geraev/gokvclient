# gokvclient

Интерактивный cli клиент для доступа к сервису GoCache. 

###Запуск программы
При запуске программа попадает в интерактивный режим.

В при запуске можно указать:

    -s или --host: хост и порт сервиса (пример: redis.com:8080). Значение по умолчанию localhost:8081
    -u или --username: имя пользователя
    -p или --password: пароль пользователя    
    -h или --help: справка по параметрам программы

Так же эти параметры можно задать в интерактивном режиме

####Пример запуска
```shell script
gokvclient --host localhost:8081 -u iqoption -p qwerty64
```

###Интерактивный режим
Интерактивный режим реализует основной функционал доступа к сервису кеша.

Команда help выдаст информацию по всем командам
```shell script
GoCache interactive client
>>> help

Commands:
  clear       clear the screen
  exit        exit the program
  help        display help
  host        set hostname and port
  key         get value for key (and internal key)
  keys        get all keys in cache
  login       set username and password for http basic authentication on endpoint
  remove      remove key
  set         set or update value
```

Для развернутой справки по комманде введите <название комманды> help
Пример:
```shell script
>>> set help

Set or update value
Examples:
  set string new_key '{"value": "string_value", "ttl": 10000}'
  set list planets '{"value": ["earth","jupiter","saturn"], "ttl": 10000}'
  set dictionary planets_map '{"value": ["earth":2220,"jupiter":3899,"saturn":23000], "ttl": 10000}'

```