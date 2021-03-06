# Документация к microshars config

# Цель

Создать комплекс решений для встраивания запущенных локально сервисов в удаленный кластер.

# Описание

Комплекс представляет из себя сервис, хранящий конфигурацию других сервисов, а так же набор библиотек для других языков, представляющих из себя реализацию Service locator для внутренних и внешних сервисов, а так же, связанных с сервисом microshards config и управляемых им.

Решение позволит встраивать или подменять локальными сервисами развернутые приложения с помощью vpn/туннелирования/сервисов проксирования локальных портов в интернет.


# Мотивация

При разработке приложений, состоящих из нескольких приложений возникают некоторые проблемы с DX (developer experiense), вот, некоторые из них:
1. Нужно поднимать весь кластер kubernets / большой docker-compose локально, чтоб проверить, что 1 микросервис работает правильно.
2. Настраивание hotreload для всех сервисов в проекте для development режима может оказаться очень большой работой
3. Может возникнуть необходимость проверить как микросервис на текущей своей стадии может интегрировать в production окружении кластера (stage стенд).
4. Некоторые вещи дебажатся только в prod/stage окружении.

# Http endpoints

## As server

GET    /

GET    /api/users
POST   /api/users
GET    /api/users/:username
PATCH  /api/users/:username
DELETE /api/users/:username

POST   /api/auth/login
POST   /api/auth/reset_password

POST   /state/refresh

GET    /api/configs
GET    /api/configs/my
POST   /api/configs
GET    /api/configs/:id ```{public: [{name: 'some name', value: 'some value'}], services: [{name: 'service_name', values: []}]}```
PATCH  /api/configs/:id
DELETE /api/configs/:id

# GRPC actions

## As server

GetConfig{ code }
# NATS actions as publisher

FetchMeta
SetConfig