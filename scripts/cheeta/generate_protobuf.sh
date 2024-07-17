#!/bin/bash

SRV_NAME="$1"
PROJECT_NAME=$(grep 'module' ../../go.mod | awk -F ' ' '{print $2}')
DIRECTORY="../../protobuf/$SRV_NAME"
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

cat << EOF > $DIRECTORY/events_v1.proto
syntax = "proto3";

package $SRV_NAME;

option go_package = "$PROJECT_NAME/internal/$SRV_NAME/pb";

message Events_V1 {
    message Tested {
        message Attributes {
            string msg = 1;
        }
        string id = 1;
        int64 occurred_on = 2;
        Attributes attributes = 3;
    }
}
EOF

exit 0
