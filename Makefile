COMPOSE_PATH=crud-service/deploy/docker-compose.yml

run-kube:
	minikube start

deploy-pg: run-kube
	helm upgrade --install my-postgres ${PWD}/k8s/helm/postgres

deploy-service:
	helm upgrade --install my-app  ${PWD}/k8s/helm/user-service

up: 
	docker compose -f $(COMPOSE_PATH) up -d --build --force-recreate

down:
	docker compose -f $(COMPOSE_PATH) down

.PHONY: up down run-kube deploy-pg deploy-service