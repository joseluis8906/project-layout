#!/bin/bash

SRV_NAME="$1"
DIRECTORY="../$SRV_NAME"
if [ ! -d "$DIRECTORY" ]; then
    mkdir -p "$DIRECTORY"
    mkdir -p "$DIRECTORY/reqs"
fi

cat << EOF > $DIRECTORY/Makefile
PROJECT_ROOT="../.."
PROJECT_NAME="$SRV_NAME"

.PHONY: build
build:
	@go build -o \$(PROJECT_ROOT)/bin/release/\$(PROJECT_NAME) \$(PROJECT_ROOT)/cmd/\$(PROJECT_NAME)

.PHONY: run
run:
	@cd \$(PROJECT_ROOT) && CONFIG_URL="http://$INFRA_HOST:2379" OTEL_SERVICE_NAME="$SRV_NAME" ./bin/release/\$(PROJECT_NAME)

.PHONY: debug
debug:
	@go build -gcflags=all="-N -l" -o \$(PROJECT_ROOT)/bin/debug/\$(PROJECT_NAME) \$(PROJECT_ROOT)/cmd/\$(PROJECT_NAME)/main.go
	@cd \$(PROJECT_ROOT) && CONFIG_URL="http://layout.local:2379" OTEL_SERVICE_NAME="prueba" dlv exec --headless --listen=:2345 --api-version=2 -- ./bin/debug/\$(PROJECT_NAME)
EOF

cat << EOF >> ../Makefile
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

cat << EOF > $DIRECTORY/reqs/req.sh
#!/bin/bash

SERVICE=\$(cat \$1 | yq '.service')
PROTO=\$(cat \$1 | yq '.proto')
METHOD=\$(cat \$1 | yq '.method')
HEADERS=\$(cat \$1 | yq '.headers')
BODY=\$(cat \$1 | yq '.body' -o=json)

grpcurl -import-path=../../../protobuf -import-path ../../../protobuf/\$SERVICE -proto "\$PROTO" -H "\$HEADERS" -d "\$BODY" -plaintext localhost:50051 \$METHOD | jq 
EOF

chmod +x $DIRECTORY/reqs/req.sh

cat << EOF > $DIRECTORY/reqs/hello_world.yml
service: $SRV_NAME 
proto: hello.proto
method: $SRV_NAME.HelloService.World
headers:
body:
  msg: it works
EOF

exit 0
