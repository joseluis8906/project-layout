#!/bin/bash

etcdctl --endpoints=$INFRA_HOST:2379 del /configs/$1 
