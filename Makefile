dev-backend-run:
	docker compose down && \
	docker compose up redis mariadb -d && \
	cd backend && go run main.go run --port 8082

dev-frontend-run:
	cd frontend && pnpm run dev