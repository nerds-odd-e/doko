
include .env
export

start:
	@echo "=== Start TDD Service ==="
	docker-compose up --build

down:
	@docker stop $(DB_HOST)
	@docker-compose --log-level ERROR down

database:
	@docker start $(DB_HOST) || \
	docker run \
	--name $(DB_HOST) \
	--net tdd \
	-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
	-p 5432:5432 \
	-d postgres
	@sleep 3
	@make main_database
	@make test_database
	@make migration

main_database:
	@docker exec \
	-it $(DB_HOST) createdb \
	-U postgres $(DB_NAME)
	@sleep 3

test_database:
	@docker exec \
	-it $(DB_HOST) createdb \
	-U postgres $(DB_NAME)_test 
	@sleep 3

migration:
	@echo "=== Run migrations and insert seed for main database ==="
	@cd migrations && DB_HOST=localhost go run *.go init;
	@cd migrations && DB_HOST=localhost go run *.go up;
	@cd migrations && DB_HOST=localhost go run *.go seed;
	@echo "=== Run migrations for test database ==="
	@cd migrations && DB_HOST=localhost DB_NAME=$(DB_NAME)_test go run *.go init;
	@cd migrations && DB_HOST=localhost DB_NAME=$(DB_NAME)_test go run *.go up;

truncate:
	@echo "=== Run truncate for main database ==="
	@cd migrations && DB_HOST=localhost go run *.go truncate;
	@echo "=== Run truncate for test database ==="
	@cd migrations && DB_HOST=localhost DB_NAME=$(DB_NAME)_test go run *.go truncate;

test:
	@DB_NAME=$(DB_NAME)_test \
	DB_HOST=localhost \
	DB_USER=$(DB_USER) \
	DB_PASSWORD=$(DB_PASSWORD) \
	go test ./... -cover -count=1 -p 1 -v -race


network:
	@docker network create -d bridge tdd
	