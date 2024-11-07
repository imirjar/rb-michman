# Michman App

Hi! My name is Michman and I can manage all of your databases in a single API. I consist of three parts:

>**Diver** is a database agent that must be placed in a database local network or in a safe access point.

>**Michman** is a single API that can manage all of existing divers.

>**AUTH** you cam use https://github.com/imirjar/rb-auth for internal auth

# INSTRUCTION

First of all you must input ./.env file in the root directory of the project.

## Michman
Подготовка к запуску проекта

    Записываем параметры в "config/config.yml" для локального запуска
    Записываем .env файл в корне директории переменные

Запустить в боевом режиме (фоном в docker)

Здесь параметры из .env перемещаются в переменные окружения в контейнере докер. Их приоритет выше чем параметры из "config/config.json"

make prod

Запускаем проект в консоли (если хочется посмотреть логи)

.env тут не будет использоваться

make start


