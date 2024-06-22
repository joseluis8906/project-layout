#!/bin/bash
#

SRV_NAME="$1"

echo "📢 following directories and files are gonna be deleted:"
echo "- cmd/$SRV_NAME"
echo "- configs/$SRV_NAME.yml"
echo "- internal/$SRV_NAME"
echo "- protobuf/$SRV_NAME"
echo "- scripts/$SRV_NAME"
echo "- scripts/reqs/$SRV_NAME"
echo -n "🔔 this operation can not be undone, do you wanna procceed? Y/n: "
read ans;

if [[ "${ans,,}" != "y" ]]; then
    exit 0;
fi

DIRECTORY="../cmd/$SRV_NAME"
echo "🔥 deleting $DIRECTORY"
if [ -d "$DIRECTORY" ]; then
    rm -rf $DIRECTORY
fi

DIRECTORY="../configs/$SRV_NAME.yml"
echo "🔥 deleting $DIRECTORY"
if [ -f "$DIRECTORY" ]; then
    rm $DIRECTORY
fi

DIRECTORY="../internal/$SRV_NAME"
echo "🔥 deleting $DIRECTORY"
if [ -d "$DIRECTORY" ]; then
    rm -rf $DIRECTORY
fi

DIRECTORY="../protobuf/$SRV_NAME"
echo "🔥 deleting $DIRECTORY"
if [ -d "$DIRECTORY" ]; then
    rm -rf $DIRECTORY
fi

DIRECTORY="../scripts/$SRV_NAME"
echo "🔥 deleting $DIRECTORY"
if [ -d "$DIRECTORY" ]; then
    rm -rf $DIRECTORY
fi

DIRECTORY="../scripts/reqs/$SRV_NAME"
echo "🔥 deleting $DIRECTORY"
if [ -d "$DIRECTORY" ]; then
    rm -rf $DIRECTORY
fi

sed -i "/# $SRV_NAME/d" Makefile
sed -i "/.PHONY: $SRV_NAME-build/d" Makefile
sed -i "/$SRV_NAME-build:/d" Makefile
sed -i "/@cd scripts\/$SRV_NAME && make build/d" Makefile
sed -i "/.PHONY: $SRV_NAME-run/d" Makefile
sed -i "/$SRV_NAME-run:/d" Makefile
sed -i "/@cd scripts\/$SRV_NAME && make run/d" Makefile
sed -i "/.PHONY: $SRV_NAME-debug/d" Makefile
sed -i "/$SRV_NAME-debug:/d" Makefile
sed -i "/@cd scripts\/$SRV_NAME && make debug/d" Makefile
sed -i '/^$/{N;/^\n$/d;}' Makefile
echo "" >> Makefile
