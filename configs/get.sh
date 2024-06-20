#!/bin/bash
#
etcdctl --endpoints=yummies.local:2379 get /configs/$1 --print-value-only | yq
