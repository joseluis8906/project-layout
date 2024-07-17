#!/bin/bash

function create() {
    echo "📢 following directories and files are gonna be created:"
    echo "+ cmd/$SRV_NAME"
    echo "+ configs/$SRV_NAME.yml"
    echo "+ internal/$SRV_NAME"
    echo "+ protobuf/$SRV_NAME"
    echo "+ scripts/$SRV_NAME"

    echo -n "🔔 this operation can not be undone, do you wanna procceed? Y/n: "
    read ans;

    if [[ "${ans,,}" != "y" ]]; then
        exit 0;
    fi

    echo "🐄 generating cmd"
    ./generate_cmd.sh $SRV_NAME

    echo "🐅 generating config"
    ./generate_config.sh $SRV_NAME

    echo "🐉 generating project"
    ./generate_project.sh $SRV_NAME

    echo "🐋 generating protocol buffer"
    ./generate_protobuf.sh $SRV_NAME

    echo "🐒 generating scripts"
    ./generate_scripts.sh $SRV_NAME

    cd ../../ && make pb 
    cd ./configs/ && ./put.sh "$SRV_NAME.yml" >> /dev/null && cd ../scripts/cheeta
}

function delete() {
    echo "📢 following directories and files are gonna be deleted:"
    echo "- cmd/$SRV_NAME"
    echo "- configs/$SRV_NAME.yml"
    echo "- internal/$SRV_NAME"
    echo "- protobuf/$SRV_NAME"
    echo "- scripts/$SRV_NAME"
    echo -n "🔔 this operation can not be undone, do you wanna procceed? Y/n: "
    read ans;

    if [[ "${ans,,}" != "y" ]]; then
        exit 0;
    fi

    cd ../../configs/ && ./del.sh "$SRV_NAME.yml" >> /dev/null && cd ../scripts/cheeta

    DIRECTORY="../../cmd/$SRV_NAME"
    echo "🔥 deleting $DIRECTORY"
    if [ -d "$DIRECTORY" ]; then
        rm -rf $DIRECTORY
    fi

    DIRECTORY="../../configs/$SRV_NAME.yml"
    echo "🔥 deleting $DIRECTORY"
    if [ -f "$DIRECTORY" ]; then
        rm $DIRECTORY
    fi

    DIRECTORY="../../internal/$SRV_NAME"
    echo "🔥 deleting $DIRECTORY"
    if [ -d "$DIRECTORY" ]; then
        rm -rf $DIRECTORY
    fi

    DIRECTORY="../../protobuf/$SRV_NAME"
    echo "🔥 deleting $DIRECTORY"
    if [ -d "$DIRECTORY" ]; then
        rm -rf $DIRECTORY
    fi

    DIRECTORY="../../scripts/$SRV_NAME"
    echo "🔥 deleting $DIRECTORY"
    if [ -d "$DIRECTORY" ]; then
        rm -rf $DIRECTORY
    fi

    FILE="../../bin/debug/$SRV_NAME"
    if [ -f "$FILE" ]; then
        rm $FILE
    fi

    FILE="../../bin/release/$SRV_NAME"
    if [ -f "$FILE" ]; then
        rm $FILE
    fi

    sed -i "/# $SRV_NAME/d" ../Makefile
    sed -i "/.PHONY: $SRV_NAME-build/d" ../Makefile
    sed -i "/$SRV_NAME-build:/d" ../Makefile
    sed -i "/@cd scripts\/$SRV_NAME && make build/d" ../Makefile
    sed -i "/.PHONY: $SRV_NAME-run/d" ../Makefile
    sed -i "/$SRV_NAME-run:/d" ../Makefile
    sed -i "/@cd scripts\/$SRV_NAME && make run/d" ../Makefile
    sed -i "/.PHONY: $SRV_NAME-debug/d" ../Makefile
    sed -i "/$SRV_NAME-debug:/d" ../Makefile
    sed -i "/@cd scripts\/$SRV_NAME && make debug/d" ../Makefile
    sed -i '/^$/{N;/^\n$/d;}' ../Makefile
    echo "" >> ../Makefile
    gawk -i inplace '/./ { e=0 } /^$/ { e += 1 } e <= 1' ../Makefile
}

function usage() {
    printf "cheeta is a helper script to create and delete services in a secure way.\n\n"
    printf "\t-c name\n"
    printf "\t\tcreates a set of directories and files with =GO code and bashscript code.\n\n"
    printf "\t-d name\n"
    printf "\t\tremoves all the files related to the provided service name.\n"
    printf "example:\n"
    printf "./cheeta -c prueba\n"
    printf "./cheeta -d prueba\n"
    exit 1
}

if [ "$#" -eq 0 ]; then
    usage
fi

while getopts 'c:d:' flag; do
    case "${flag}" in
        c) SRV_NAME="${OPTARG}" create;;
        d) SRV_NAME="${OPTARG}" delete;;
        *) usage;;
    esac
done
