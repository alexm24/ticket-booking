# Rest api for Ticket booking

## Postgres

#### create docker postgres
- docker run --name=postgres -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d postgres

#### create init database
- migrate create -ext sql -dir ./migrations -seq init

#### create migration
https://github.com/golang-migrate/migrate
- migrate -path ./migrations -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" up

## Swagger
- https://github.com/alexm24/ticket-booking/blob/main/internal/handler/api/api.swagger.yaml