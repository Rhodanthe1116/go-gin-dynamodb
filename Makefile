NAME=covid-tracker-server
VERSION=0.0.1

.PHONY: build
## build: Compile the packages.
build:
	@go build -o $(NAME)

.PHONY: run
## run: Build and Run in development mode.
run: build
	@./$(NAME) -e development -s all

.PHONY: user
## user: Build and Run in development mode with user service only.
user: build
	@./$(NAME) -e development -s user

.PHONY: store
## store: Build and Run in development mode with store service only.
store: build
	@./$(NAME) -e development -s store

.PHONY: record
## record: Build and Run in development mode with record service only.
record: build
	@./$(NAME) -e development -s record

.PHONY: docker
## docker: Run docker-compose.
docker:
	@(cd .devcontainer; docker ps | grep devcontainer_web_1 || docker-compose up -d;)
	docker exec -it devcontainer_web_1 zsh

.PHONY: dbenv
## dbenv: export AWS credentials
dbenv:


.PHONY: stop
## stop: Stop docker-compose.
stop:
	@(cd .devcontainer; docker-compose down;)

.PHONY: run-prod
## run-prod: Build and Run in production mode.
run-prod: build
	@./$(NAME) -e production

.PHONY: clean
## clean: Clean project and previous builds.
clean:
	@rm -f $(NAME)

.PHONY: deps
## deps: Download modules
deps:
	@go mod download

.PHONY: test
## test: Run tests with verbose mode
test:
	@go test -v ./tests/*

.PHONY: help
all: help
# help: show this help message
help: Makefile
	@echo
	@echo " Choose a command to run in "$(APP_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
