#!/bin/bash
#

SRV_NAME="$1"
PROJECT_NAME=$(grep 'module' ../go.mod | awk -F ' ' '{print $2}')
DIRECTORY="../protobuf/$SRV_NAME"
if [ ! -d "$DIRECTORY" ]; then
    mkdir -p "$DIRECTORY"
fi

cat << EOF > $DIRECTORY/hello.proto
syntax = "proto3";

package $SRV_NAME;

option go_package = "$PROJECT_NAME/internal/$SRV_NAME/pb";

service HelloService {
    rpc World(HelloWorldRequest) returns (HelloWorldResponse);
}

message HelloWorldRequest {
    string msg = 1;
}

message HelloWorldResponse {
    string msg = 1;
}
EOF

exit 0
