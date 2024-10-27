#!/bin/bash

SERVICE=$(cat $1 | yq '.service')
PROTO=$(cat $1 | yq '.proto')
METHOD=$(cat $1 | yq '.method')
HEADERS=$(cat $1 | yq '.headers')
BODY=$(cat $1 | yq '.body' -o=json)

grpcurl -import-path=../../../protobuf -import-path ../../../protobuf/$SERVICE -proto "$PROTO" -H "$HEADERS" -d "$BODY" -plaintext localhost:50051 "$SERVICE.$METHOD" | jq 
