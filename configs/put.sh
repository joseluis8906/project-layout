#!/bin/bash

etcdctl --endpoints=$INFRA_HOST:2379 put /configs/$1 "$(cat -v ./$1)"
