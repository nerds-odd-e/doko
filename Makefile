include .env
export

test:
	@DB_NAME=$(DB_NAME)_test \
	DB_HOST=localhost \
	DB_USER=$(DB_USER) \
	DB_PASSWORD=$(DB_PASSWORD) \
	go test ./... -cover -count=1 -p 1 -v -race

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

tcr:
	make test || (git reset --hard; echo "----- TCR reverted -----"; exit 0)
	git add .
	git commit -am "tcring" | tee | grep "working tree clean$$" && (echo "----- TCR passed -----"; exit 0)
	echo "----- TCR no change found -----"

limbo: tcr
	git pull --rebase | grep -E "u1p to date\.$$" && git push || echo "No changes"
