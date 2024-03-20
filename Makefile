include .env

export DATABASE_URL ?= postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)
export UNIT_TEST_COVERAGE ?= 60.0

bin:
	@mkdir -p bin

setup-tools: bin
	@curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
ifeq ($(shell uname), Linux)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/ 
else ifeq ($(shell uname), Darwin)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.darwin-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/ 
else
	@echo "Your OS is not supported."
endif

migration-create:
	bin/migrate create -dir migrations -tz UTC -ext sql $(name)

migration-up:
	bin/migrate -path migrations -database "${DATABASE_URL}" up

migration-down:
	bin/migrate -path migrations -database "${DATABASE_URL}" down $(n)

run-dev:
	bin/air -d --build.cmd "go build -o tmp/api main.go" --build.bin "tmp/api"

run:
	./main

build:
	go build ./main.go
