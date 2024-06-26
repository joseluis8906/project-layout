#!/bin/bash
#
project=$(grep 'module' ./go.mod | awk -F ' ' '{print $2}')
protos=$(fd ".proto" ./protobuf/**/)
for proto in $protos
do
    pbpath="$(dirname ${proto/protobuf/internal})/pb"
    true || rm "$pbpath/*.go"
    module="$project/$(dirname ${pbpath/.\//})/pb"
    protoc --proto_path=protobuf\
		--go_out="$pbpath" --go_opt=module="$module"\
		--go-grpc_out="$pbpath" --go-grpc_opt=module="$module"\
	    $proto
done

protos=$(ls ./protobuf -1 | rg .proto)
for proto in $protos
do
    pbpath="./pkg/pb"
    true || rm "$pbpath/*.go"
    module="$project/pkg/pb"
    protoc --proto_path=protobuf\
		--go_out="$pbpath" --go_opt=module="$module"\
	    $proto	
done
