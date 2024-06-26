#!/bin/bash
#

SRV_NAME="$1"
DIRECTORY="../configs"
cat << EOF > "$DIRECTORY/$SRV_NAME.yml"
app:
  name: $SRV_NAME
otel:
  endpoint: http://yummies.local:4317
fluentd:
  host: yummies.local
  port: 24224
  network: tcp
  json: true
  tag: my.logs 
grpc:
  port: 50051
http:
  addr: ':9090'
nats:
  url: 'nats://yummies.local:4222'
kafka:
  bootstrap:
    servers: 'yummies.local:9092'
  acks: all
  group:
    id: $SRV_NAME
  auto:
    offset:
      reset: earliest
  topics:
    - v1.tested
mongodb:
  uri: 'mongodb://yummies:yummies@yummies.local:27017/$SRV_NAME?authMechanism=PLAIN'
rabbitmq:
    url: 'amqp://guest:guest@yummies.local:5672/'
pprof: true
EOF

exit 0
