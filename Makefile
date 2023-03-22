start:
	docker compose -f deployments/docker-compose.dev.yml up -d

start_server:
	docker compose -f deployments/docker-compose.dev.yml up -d db server

start_db:
	docker compose -f deployments/docker-compose.dev.yml up -d db

start_client: 
	docker compose -f deployments/docker-compose.dev.yml up -d client

logsf:
	docker compose -f deployments/docker-compose.dev.yml logs --tail 100 -f

logs:
	docker compose -f deployments/docker-compose.dev.yml logs --tail 100

down:
	docker compose -f deployments/docker-compose.dev.yml down --remove-orphans