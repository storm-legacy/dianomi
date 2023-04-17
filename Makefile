PWD=$(shell pwd)

init: generate_ssl

start:
	docker compose -f deployments/docker-compose.dev.yml up -d

startf:
	docker compose -f deployments/docker-compose.dev.yml up

start_proxy:
	docker compose -f deployments/docker-compose.dev.yml up -d db haproxy

start_server:
	docker compose -f deployments/docker-compose.dev.yml up -d db server

start_db:
	docker compose -f deployments/docker-compose.dev.yml up -d db pgadmin

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

generate_ssl:
	if ! [ -f "$(PWD)/deployments/ssl/self-signed.crt" ]; then \
		docker run --rm --name omgwtfssl \
		-e CA_SUBJECT="snake-oil" \
		-e CA_EXPIRE="3650" \
		-e SSL_SIZE="2048" \
		-e SSL_EXPIRE="3650" \
		-e SSL_SUBJECT="localhost" \
		-e SSL_IP="127.0.0.1" \
		-v "$(PWD)/deployments/ssl/:/certs:rw" \
		--user 1000:1000 \
		paulczar/omgwtfssl:latest; \
		cat "$(PWD)/deployments/ssl/key.pem" "$(PWD)/deployments/ssl/cert.pem" > "$(PWD)/deployments/ssl/self-signed.crt"; \
	fi