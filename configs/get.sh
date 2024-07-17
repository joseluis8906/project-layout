#!/bin/bash
#
etcdctl --endpoints=$INFRA_HOST:2379 get /configs/$1 --print-value-only | yq
