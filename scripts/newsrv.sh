#!/bin/bash
#

SRV_NAME="$1"

echo "🐄 generating cmd"
./newsrv/generate_cmd.sh $SRV_NAME

echo "🐅 generating config"
./newsrv/generate_config.sh $SRV_NAME

echo "🐉 generating project"
./newsrv/generate_project.sh $SRV_NAME

echo "🐋 generating protocol buffer"
./newsrv/generate_protobuf.sh $SRV_NAME

echo "🐒 generating scripts"
./newsrv/generate_scripts.sh $SRV_NAME

