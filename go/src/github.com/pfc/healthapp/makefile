.PHONY: all dev clean env-up env-down run

all: clean env-up run

dev: run

##### ENV
env-up:
	@echo "Start environment ..."
	@docker-compose up --force-recreate -d
	@echo "Environment up"

env-down:
	@echo "Stop environment ..."
	@docker-compose down
	@echo "Environment down"

##### RUN
run:
	@echo "Start app ..."
	@./main

##### CLEAN
clean: env-down
	@echo "Clean up ..."
	@rm -rf /tmp/health-app* heroes-service
	@sudo docker rm -f -v `docker ps -a --no-trunc | grep "healthapp" | cut -d ' ' -f 1` 2>/dev/null || true
	@sudo docker rmi `docker images --no-trunc | grep "healthapp" | cut -d ' ' -f 1` 2>/dev/null || true
	@echo "Clean up done"
