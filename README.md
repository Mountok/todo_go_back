### Разворачиваем potgres в докер

1. команда для скачивания пострес образа:
```docker pull postgres```
2. разворачиваем контейнер бд:
```docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres```
3. для просмотра контейнеров которые запущены:
```docker ps```
4. создание файлов миграции:
``` migrate create -ext sql -dir ./schema -seq init```
5.  миграция: 
```migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up ```