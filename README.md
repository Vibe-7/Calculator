    Calculator

**Описание проекта:**
В этом проекте я написал работающий **Backend сервер** на **Go**, реализовал **архитектуру** с разделением **слоёв**, подключил современный **веб-фреймворк Echo**, а также несколько вспомогательных **пакетов** для **обработки выражений** и **работы с БД**.

Backend взаимодействует с базой данных **PostgreSQL**, где хранятся расчёты. Для развертывания всей инфраструктуры использовал **Docker**, который объединяет **Backend**, **Frontend** (скачанный шаблон на Node.js) и **PostgreSQL** в единую работающую систему.

**Стек технологий:**
**Go** — язык программирования.
**Echo** — фреймворк для создания REST API.
**Gorm** — ORM для работы с PostgreSQL.
**Govaluate** — пакет для вычисления математических выражений.
**PostgreSQL** — реляционная база данных.
**Docker** — для контейнеризации Backend, Frontend и БД.
**Node.js** — Frontend часть (готовый шаблон).

**Архитектура:**
**Handlers** — обработчики HTTP-запросов.
**Services** — бизнес-логика.
**Repository** — слой работы с базой данных (PostgreSQL через Gorm).
**Модели** — описание структуры расчётов.
**Валидация** — базовая проверка данных на этапе получения запросов.
