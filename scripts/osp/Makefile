PROJECT_ROOT="../.."
PROJECT_NAME="osp"

.PHONY: build
build:
	@go build -o $(PROJECT_ROOT)/bin/release/$(PROJECT_NAME) $(PROJECT_ROOT)/cmd/$(PROJECT_NAME)

.PHONY: run
run:
	@CONFIG_URL="http://yummies.local:2379" OTEL_SERVICE_NAME="osp" $(PROJECT_ROOT)/bin/release/$(PROJECT_NAME)

.PHONY: debug
debug:
	@go build -gcflags=all="-N -l" -o $(PROJECT_ROOT)/bin/debug/$(PROJECT_NAME) $(PROJECT_ROOT)/cmd/$(PROJECT_NAME)/main.go
	@CONFIG_URL="http://yummies.local:2379" OTEL_SERVICE_NAME="osp" dlv exec --headless --listen=:2345 --api-version=2 -- $(PROJECT_ROOT)/bin/debug/$(PROJECT_NAME)
