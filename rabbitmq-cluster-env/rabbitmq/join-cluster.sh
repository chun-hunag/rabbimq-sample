#!/bin/bash
# start rabbitmq service in backgorund (erlang node and rabbitmq application , detached mean in background)
rabbitmq-server -detached
sleep 5s

# stop rabbitmq application
rabbitmqctl stop_app

# join rabbimtq cluster
rabbitmqctl join_cluster ${MASTER_NODE_HOST_NAME}

# start rabbitmq application 來 sync 資料
rabbitmqctl start_app

# stop rabbitmq service (errlang node and rabbitmq application )
rabbitmqctl stop
sleep 5s

# start rabbitmq service foreground
rabbitmq-server