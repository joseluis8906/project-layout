#!/bin/bash
#
etcdctl --endpoints=yummies.local:2379 del /configs/$1 
