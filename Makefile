# Copyright (c) 2023 Julian Klaiber

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: generate-proto update-submodules

install-deps: ## Install development dependencies
	sudo apt install -y protobuf-compiler
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

update-submodules: ## Update submodules
	git submodule update --remote --merge

generate-proto: ## Generate Go code from proto files
	protoc --go_out=. --go_opt=Mproto/intent.proto=pkg/api --go-grpc_out=. --go-grpc_opt=Mproto/intent.proto=pkg/api proto/*.proto

go-run: ## Run the server with go run
	go run dc/main.go server --config test_asset/config.yaml

build: ## Build the server
	go build -o out/bin/dummy-controller dc/main.go

run: ## Run the server
	./out/bin/dummy-controller server --config test_asset/config.yaml

help: ## Show this help message
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_0-9-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)