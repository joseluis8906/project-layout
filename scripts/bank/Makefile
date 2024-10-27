PROJECT_ROOT="../.."
PROJECT_NAME="bank"

.PHONY: build
build:
	@go build -o $(PROJECT_ROOT)/bin/release/$(PROJECT_NAME) $(PROJECT_ROOT)/cmd/$(PROJECT_NAME)

.PHONY: run
run:
	@cd $(PROJECT_ROOT) && CONFIG_URL="http://layout.local:2379" OTEL_SERVICE_NAME="bank" ./bin/release/$(PROJECT_NAME)

.PHONY: debug
debug:
	@go build -gcflags=all="-N -l" -o $(PROJECT_ROOT)/bin/debug/$(PROJECT_NAME) $(PROJECT_ROOT)/cmd/$(PROJECT_NAME)/main.go
	@cd $(PROJECT_ROOT) && CONFIG_URL="http://layout.local:2379" OTEL_SERVICE_NAME="prueba" dlv exec --headless --listen=:2345 --api-version=2 -- ./bin/debug/$(PROJECT_NAME)
