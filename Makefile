PWD=$(shell pwd)

start:
	docker compose -f deployments/docker-compose.dev.yml up -d client server db traefik

startf:
	docker compose -f deployments/docker-compose.dev.yml up client server db traefik

start_proxy:
	docker compose -f deployments/docker-compose.dev.yml up -d traefik

start_server:
	docker compose -f deployments/docker-compose.dev.yml up -d db server

start_db:
	docker compose -f deployments/docker-compose.dev.yml up -d db

start_pgadmin:
	docker compose -f deployments/docker-compose.dev.yml up -d pgadmin

start_client: 
	docker compose -f deployments/docker-compose.dev.yml up -d client

logsf:
	docker compose -f deployments/docker-compose.dev.yml logs --tail 100 -f

logs:
	docker compose -f deployments/docker-compose.dev.yml logs --tail 100

down:
	docker compose -f deployments/docker-compose.dev.yml down --remove-orphans

hard_down:
	docker compose -f deployments/docker-compose.dev.yml down --remove-orphans -v