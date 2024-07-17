#!/bin/bash

SRV_NAME="$1"
DIRECTORY="../../configs"
cat << EOF > "$DIRECTORY/$SRV_NAME.yml"
app:
  name: $SRV_NAME
otel:
  endpoint: http://$INFRA_HOST:4317
fluentd:
  host: $INFRA_HOST
  port: 24224
  network: tcp
  json: true
  tag: my.logs 
grpc:
  port: 50051
http:
  addr: ':9191'
nats:
  url: 'nats://$INFRA_HOST:4222'
kafka:
  bootstrap:
    servers: '$INFRA_HOST:9093'
  acks: all
  group:
    id: $SRV_NAME
  auto:
    offset:
      reset: earliest
  topics:
    - v1.tested
mongodb:
  uri: 'mongodb://$INFRA_USER:$INFRA_PASSWD@$INFRA_HOST:27017/$SRV_NAME?authMechanism=PLAIN'
rabbitmq:
    url: 'amqp://$INFRA_USER:$INFRA_PASSWD@$INFRA_HOST:5672/'
pprof: true
EOF

exit 0
