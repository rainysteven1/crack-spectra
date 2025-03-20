dev-run:
	docker compose down && \
	docker compose up redis mariadb -d && \
	cd backend && go run main.go run --port 8082