#!/bin/bash
#
etcdctl --endpoints=yummies.local:2379 put /configs/$1 "$(cat -v ./$1)"
