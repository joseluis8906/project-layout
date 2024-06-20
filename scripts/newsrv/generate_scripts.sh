#!/bin/bash
#

SRV_NAME="$1"
DIRECTORY="./$SRV_NAME"
if [ ! -d "$DIRECTORY" ]; then
    mkdir -p "$DIRECTORY"
fi

cat << EOF > $DIRECTORY/Makefile
PROJECT_ROOT="../.."
PROJECT_NAME="$SRV_NAME"

.PHONY: build
build:
	@go build -o \$(PROJECT_ROOT)/bin/release/\$(PROJECT_NAME) \$(PROJECT_ROOT)/cmd/\$(PROJECT_NAME)

.PHONY: run
run:
	@CONFIG_URL="http://yummies.local:2379" OTEL_SERVICE_NAME="$SRV_NAME" \$(PROJECT_ROOT)/bin/release/\$(PROJECT_NAME)

.PHONY: debug
debug:
	@go build -gcflags=all="-N -l" -o \$(PROJECT_ROOT)/bin/debug/\$(PROJECT_NAME) \$(PROJECT_ROOT)/cmd/\$(PROJECT_NAME)/main.go
	@CONFIG_URL="http://yummies.local:2379" OTEL_SERVICE_NAME="$SRV_NAME" dlv exec --headless --listen=:2345 --api-version=2 -- \$(PROJECT_ROOT)/bin/debug/\$(PROJECT_NAME)
EOF

cat << EOF >> Makefile
# $SRV_NAME
.PHONY: $SRV_NAME-build
$SRV_NAME-build:
	@cd scripts/$SRV_NAME && make build

.PHONY: $SRV_NAME-run
$SRV_NAME-run:
	@cd scripts/$SRV_NAME && make run

.PHONY: $SRV_NAME-debug
$SRV_NAME-debug:
	@cd scripts/$SRV_NAME && make debug

EOF

DIRECTORY="./reqs/$SRV_NAME"
if [ ! -d "$DIRECTORY" ]; then
    mkdir -p "$DIRECTORY"
fi

cat << EOF > $DIRECTORY/hello_world.yml
service: $SRV_NAME 
proto: hello.proto
method: $SRV_NAME.HelloService.World
headers:
body:
  msg: it works
EOF

exit 0
