DB_DSN := "postgres://postgres:abdul@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

clear:
	migrate -path ./migrations -database "postgres://postgres:abdul@localhost:5432/postgres?sslmode=disable" drop -f

run:
	go run cmd/app/main.go

create-container:
	docker run --name container -e POSTGRES_PASSWORD=abdul -d -p 5432:5432 postgres

lint:
	golangci-lint run --color=auto

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

test:
	go test ./... -v
