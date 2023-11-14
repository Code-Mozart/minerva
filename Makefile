stop-local:
	@-pkill "minerva-backend"

run-local: stop-local
	@cd ./minerva-backend && make run &
	@cd ./minerva-webapp && make run

run-docker:
	@echo "Running Docker Compose"
	@docker-compose up -d
