DOCKERCOMPOSECMD=docker-compose
GOCMD=go

dc-up: $(DOCKERCOMPOSECMD) up -d --build

dc-down: $(DOCKERCOMPOSECMD) down --remove-orphans

dc-restart: dc-down dc-up

fmt: $(GOCMD) fmt ./...

test-clean: fmt $(GOCMD) clean -testcache

tests: fmt test-clean $(GOCMD) test -cover -p=1 ./...